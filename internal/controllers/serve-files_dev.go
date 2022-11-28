//go:build dev

package controllers

import (
	"net/http"
	"path/filepath"
)

var (
	DistDir = filepath.Join("views", "dist")
)

func ServeFiles(pattern string, mux *http.ServeMux) *http.ServeMux {
	mux.Handle(pattern, http.FileServer(http.Dir(DistDir)))
	return mux
}
