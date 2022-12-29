//go:build !dev

package fileserver

import (
	"net/http"
	"regexp"

	"github.com/knackwurstking/pirgb-web/frontend"
	"github.com/knackwurstking/pirgb-web/internal/router"
)

var (
	FS              = frontend.GetFS()
	Routes []string = []string{
		"/sections",
		"/groups",
		"/settings",
	}
)

func ServeFiles(pattern string, mux *router.RegexRouter) *router.RegexRouter {
	rePattern, _ := regexp.Compile(pattern + "(.*)")
	mux.HandleRegEx(rePattern, http.FileServer(FS), Routes...)

	return mux
}
