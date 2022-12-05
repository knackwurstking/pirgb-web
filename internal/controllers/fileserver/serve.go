//go:build !dev

package fileserver

import (
	"net/http"
	"regexp"

	frontend "github.com/knackwurstking/pirgb-web/frontend_svelte"
	"github.com/knackwurstking/pirgb-web/internal/controllers"
)

var FS = frontend.GetFS()

func ServeFiles(pattern string, mux *controllers.RegexHandler) *controllers.RegexHandler {
	mux.Handle("/", http.FileServer(FS))

	rePattern, _ := regexp.Compile(pattern + "(.*)")
	mux.HandleRegEx(rePattern, http.FileServer(FS))

	return mux
}
