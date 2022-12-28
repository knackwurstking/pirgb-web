//go:build dev

package fileserver

import (
	"net/http"
	"path/filepath"
	"regexp"

	"github.com/knackwurstking/pirgb-web/internal/router"
)

var DistDir = filepath.Join("frontend", "dist")

func ServeFiles(pattern string, mux *router.RegexRouter) *router.RegexRouter {
	rePattern, _ := regexp.Compile(pattern + "(.*)")
	mux.HandleRegEx(rePattern, http.FileServer(http.Dir(DistDir)))

	return mux
}
