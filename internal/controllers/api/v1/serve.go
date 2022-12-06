package v1

import (
	"regexp"

	"github.com/knackwurstking/pirgb-web/internal/controllers"
	"github.com/knackwurstking/pirgb-web/internal/controllers/api/v1/handler"
	"github.com/knackwurstking/pirgb-web/pkg/middleware"
)

func ServeApi(pattern string, mux *controllers.RegexHandler) *controllers.RegexHandler {
	mux.HandleFunc(pattern+"/events", middleware.Logger(handler.Events))

	regex, _ := regexp.Compile(pattern + "/devices(.*)")
	mux.HandleRegEx(regex, handler.NewDeviceHandler(pattern+"/device"))

	return mux
}
