package v1

import (
	"regexp"

	"github.com/knackwurstking/pirgb-web/internal/controllers/api/v1/handler"
	"github.com/knackwurstking/pirgb-web/pkg/middleware"
	"github.com/knackwurstking/pirgb-web/pkg/router"
)

func ServeApi(pattern string, mux *router.RegexRouter) *router.RegexRouter {
	mux.HandleFunc(pattern+"/events", middleware.Logger(handler.Events))

	regex, _ := regexp.Compile(pattern + "/devices(.*)")
	mux.HandleRegEx(regex, handler.NewDeviceHandler(pattern+"/device"))

	return mux
}
