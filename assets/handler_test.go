package assets

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// The consumer case: no assets directory next to the process, so the files
// have to come out of the embed and be cacheable.
func TestHandlerServesEmbeddedStylesheet(t *testing.T) {
	t.Setenv("GO_ENV", "")
	t.Chdir(t.TempDir())

	mux := http.NewServeMux()
	mux.Handle("GET /assets/", Handler())

	recorder := httptest.NewRecorder()
	mux.ServeHTTP(recorder, httptest.NewRequest(http.MethodGet, StylesheetURL(), nil))

	if recorder.Code != http.StatusOK {
		t.Fatalf("GET %s returned %d", StylesheetURL(), recorder.Code)
	}
	if recorder.Body.Len() == 0 {
		t.Fatal("stylesheet is empty")
	}
	if got := recorder.Header().Get("Cache-Control"); got != "public, max-age=31536000, immutable" {
		t.Fatalf("hashed URL got Cache-Control %q", got)
	}
}
