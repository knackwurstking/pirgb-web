//go:build dev

package fileserver

import (
	"net/http"
	"path/filepath"
	"regexp"

	"github.com/knackwurstking/pirgb-web/internal/router"
)

var (
	DistDir          = filepath.Join("frontend", "dist")
	Routes  []string = []string{
		"/sections",
		"/groups",
		"/settings",
	}
)

func ServeFiles(pattern string, mux *router.RegexRouter) *router.RegexRouter {
	rePattern, _ := regexp.Compile(pattern + "(.*)")
	mux.HandleRegEx(rePattern, http.FileServer(http.Dir(DistDir)), Routes...)

	return mux
}
