package assets

import (
	"crypto/sha256"
	"encoding/hex"
	"io/fs"
	"net/http"
	"os"
	"sync"
)

// Stylesheet is the compiled stylesheet shipped with the library, built from
// assets/css/globals.css by `task build-css`. It carries every component's
// styling and all eight styles; pick one with the style-<name> class on
// <body>. Projects that run Tailwind themselves can point it at
// globals.css instead and skip this file.
const Stylesheet = "css/templ-ui.css"

const assetsDir = "assets"

// diskDir is the working copy of the assets directory, or "". Same rule as the
// component script bundle: this repo's own dev server serves the files from
// disk so edits show up without a rebuild, while a project that imports the
// library has no such directory and is served the embedded copies.
func diskDir() string {
	if os.Getenv("GO_ENV") == "production" {
		return ""
	}
	if info, err := os.Stat(assetsDir); err != nil || !info.IsDir() {
		return ""
	}
	return assetsDir
}

var (
	hashOnce sync.Once
	cssHash  string
)

// StylesheetURL is the href for the shipped stylesheet, hashed over its
// contents: upgrading the library changes the URL, rebuilding the same bytes
// does not. Handler reads the hash back and caches the response forever.
//
//	<link rel="stylesheet" href={ assets.StylesheetURL() }/>
func StylesheetURL() string {
	hashOnce.Do(func() {
		css, err := Assets.ReadFile(Stylesheet)
		if err != nil {
			return
		}
		sum := sha256.Sum256(css)
		cssHash = hex.EncodeToString(sum[:8])
	})
	if cssHash == "" {
		return "/" + assetsDir + "/" + Stylesheet
	}
	return "/" + assetsDir + "/" + Stylesheet + "?v=" + cssHash
}

// Handler serves the stylesheet, fonts and images. Mount it at /assets/, the
// prefix the stylesheet's @font-face rules hardcode:
//
//	mux.Handle("GET /assets/", assets.Handler())
func Handler() http.Handler {
	return http.StripPrefix("/"+assetsDir+"/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var files http.FileSystem
		if dir := diskDir(); dir != "" {
			w.Header().Set("Cache-Control", "no-store")
			files = http.Dir(dir)
		} else {
			// A versioned URL names one immutable build; a bare one may be
			// answered by the next release, so it only gets the day.
			if r.URL.Query().Get("v") != "" {
				w.Header().Set("Cache-Control", "public, max-age=31536000, immutable")
			} else {
				w.Header().Set("Cache-Control", "public, max-age=86400")
			}
			files = http.FS(fs.FS(Assets))
		}
		http.FileServer(files).ServeHTTP(w, r)
	}))
}
