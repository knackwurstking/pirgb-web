//go:build dev

package fileserver

import (
	"net/http"
	"path/filepath"

	"github.com/knackwurstking/pirgb-web/internal/controllers"
)

var DistDir = filepath.Join("frontend_svelte", "dist")

func ServeFiles(pattern string, mux *controllers.RegExpHandler) *controllers.RegExpHandler {
	mux.Handle(pattern, http.FileServer(http.Dir(DistDir)))
	return mux
}
