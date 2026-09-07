package components

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestScriptsHandlerServesGeneratedHashedPath(t *testing.T) {
	t.Setenv("GO_ENV", "production")

	mux := http.NewServeMux()
	mux.Handle("GET /components/{bundle}", ScriptsHandler())

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, scriptsSrc(), nil)
	mux.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("GET %s returned %d", scriptsSrc(), recorder.Code)
	}
	if recorder.Body.Len() == 0 {
		t.Fatal("script bundle is empty")
	}
}

// A project that imports the library has no components directory next to it,
// and no reason to set GO_ENV=production while developing. The bundle has to
// fall back to the embedded sources instead of coming out empty.
func TestBundleFallsBackToEmbeddedSources(t *testing.T) {
	t.Setenv("GO_ENV", "")
	t.Chdir(t.TempDir())

	js, _, _, fromDisk := bundle()
	if fromDisk {
		t.Fatal("bundle was read from disk with no components directory in sight")
	}
	if !bytes.Contains(js, []byte("// components/accordion/accordion.js")) {
		t.Fatalf("embedded bundle is missing the component sources (%d bytes)", len(js))
	}
}
