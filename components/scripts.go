package components

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"encoding/hex"
	"io/fs"
	"net/http"
	"os"
	"strings"
	"sync"
)

// The component JS bundle is the concatenation of every components/*/*.js
// file. Each file is a standalone IIFE; the lexical walk order keeps
// floatingui/floating_ui_core.js ahead of floating_ui_dom.js, the one pair
// where load order matters. The bundle is built once from the embedded files;
// only this repo's own dev server rebuilds it from disk on every request, so
// edits hot-reload.

const componentsDir = "components"

// diskFS is the working copy of the component sources, or nil. It is non-nil
// only while developing this repo: the directory has to be there next to the
// running process. A project importing the library has no such directory —
// its bundle always comes from the embedded sources, which is why the check
// is for the directory and not for GO_ENV alone. Missing that, every consumer
// outside GO_ENV=production shipped an empty bundle and no error with it.
func diskFS() fs.FS {
	if os.Getenv("GO_ENV") == "production" {
		return nil
	}
	if info, err := os.Stat(componentsDir); err != nil || !info.IsDir() {
		return nil
	}
	return os.DirFS(componentsDir)
}

func buildBundle(fsys fs.FS) ([]byte, string) {
	var buf bytes.Buffer
	fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() || !strings.HasSuffix(path, ".js") || strings.HasSuffix(path, ".min.js") {
			return nil
		}
		data, err := fs.ReadFile(fsys, path)
		if err != nil {
			return nil
		}
		buf.WriteString("// components/" + path + "\n")
		buf.Write(data)
		buf.WriteString("\n")
		return nil
	})
	sum := sha256.Sum256(buf.Bytes())
	return buf.Bytes(), hex.EncodeToString(sum[:8])
}

var (
	prodOnce sync.Once
	prodJS   []byte
	prodGz   []byte
	prodHash string
)

// gzipBundle compresses once so bare Go deployments without a compressing
// proxy still ship ~5x smaller transfers; proxies that compress themselves
// simply never see the identity variant.
func gzipBundle(js []byte) []byte {
	var buf bytes.Buffer
	zw, _ := gzip.NewWriterLevel(&buf, gzip.BestCompression)
	_, _ = zw.Write(js)
	_ = zw.Close()
	return buf.Bytes()
}

// bundle returns the JS, its gzip variant (nil when rebuilt from disk) and the
// content hash. The last result reports whether the bundle came from disk, the
// one case where it must not be cached.
func bundle() ([]byte, []byte, string, bool) {
	if fsys := diskFS(); fsys != nil {
		js, hash := buildBundle(fsys)
		return js, nil, hash, true
	}
	prodOnce.Do(func() {
		prodJS, prodHash = buildBundle(TemplFiles)
		prodGz = gzipBundle(prodJS)
	})
	return prodJS, prodGz, prodHash, false
}

// The content hash lives in the path like Next's static chunks
// (/_next/static/chunks/<hash>.js): query strings are ignored by some CDN
// caches, path hashes never are.
func scriptsSrc() string {
	_, _, hash, _ := bundle()
	return "/components/shadcn-templ-" + hash + ".js"
}

// ScriptsHandler serves the component JS bundle. Mount it on
// GET /components/{bundle}: it answers the current hashed name
// (shadcn-templ-<hash>.js) and the plain shadcn-templ.js alias, 404s anything else.
func ScriptsHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		js, gz, hash, fromDisk := bundle()
		base := r.URL.Path[strings.LastIndex(r.URL.Path, "/")+1:]
		if base != "shadcn-templ.js" && base != "shadcn-templ-"+hash+".js" {
			http.NotFound(w, r)
			return
		}
		etag := `"` + hash + `"`
		w.Header().Set("Content-Type", "application/javascript")
		w.Header().Set("ETag", etag)
		if fromDisk {
			w.Header().Set("Cache-Control", "no-store")
		} else {
			w.Header().Set("Cache-Control", "public, max-age=31536000, immutable")
			if r.Header.Get("If-None-Match") == etag {
				w.WriteHeader(http.StatusNotModified)
				return
			}
		}
		if gz != nil && strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			w.Header().Set("Content-Encoding", "gzip")
			w.Header().Set("Vary", "Accept-Encoding")
			_, _ = w.Write(gz)
			return
		}
		_, _ = w.Write(js)
	})
}
