//go:build !dev

package fileserver

import (
	"net/http"
	"regexp"

	frontend "github.com/knackwurstking/pirgb-web/frontend_svelte"
	"github.com/knackwurstking/pirgb-web/pkg/router"
)

var FS = frontend.GetFS()

func ServeFiles(pattern string, mux *router.RegexRouter) *router.RegexRouter {
	rePattern, _ := regexp.Compile(pattern + "(.*)")
	mux.HandleRegEx(rePattern, http.FileServer(FS))

	return mux
}
