//go:build !dev

package fileserver

import (
	"net/http"

	frontend "github.com/knackwurstking/pirgb-web/frontend_svelte"
	"github.com/knackwurstking/pirgb-web/internal/controllers"
)

var FS = frontend.GetFS()

func ServeFiles(pattern string, mux *controllers.RegExpHandler) *controllers.RegExpHandler {
	mux.Handle("/", http.FileServer(FS))
	return mux
}
