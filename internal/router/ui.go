//go:build !dev

package router

import (
	"net/http"

	"gitlab.com/knackwurstking/pirgb-web/frontend"
)

func init() {
	Mux.Handle("/", http.FileServer(http.FS(frontend.Dist)))

	info["    GET /"] = "serve ui"
}
