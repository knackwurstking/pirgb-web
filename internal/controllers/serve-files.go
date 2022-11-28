//go:build !dev

package controllers

import (
	"net/http"

	frontent "github.com/knackwurstking/pirgb-web/views"
)

var (
	FS = frontent.GetFS()
)

func ServeFiles(pattern string, mux *http.ServeMux) *http.ServeMux {
	mux.Handle("/", http.FileServer(FS))
	return mux
}
