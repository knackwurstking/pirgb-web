//go:build !dev

package controllers

import (
	"net/http"

	frontend "github.com/knackwurstking/pirgb-web/frontend_svelte"
)

var FS = frontend.GetFS()

func ServeFiles(pattern string, mux *http.ServeMux) *http.ServeMux {
	mux.Handle("/", http.FileServer(FS))
	return mux
}
