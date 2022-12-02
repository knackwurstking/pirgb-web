//go:build dev

package fileserver

import (
	"net/http"
	"path/filepath"
)

var DistDir = filepath.Join("frontend_svelte", "dist")

func ServeFiles(pattern string, mux *http.ServeMux) *http.ServeMux {
	mux.Handle(pattern, http.FileServer(http.Dir(DistDir)))
	return mux
}
