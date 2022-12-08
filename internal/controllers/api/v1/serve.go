package v1

import (
	"regexp"

	"github.com/knackwurstking/pirgb-web/internal/controllers/api/v1/handler"
	"github.com/knackwurstking/pirgb-web/pkg/pirgb"
	"github.com/knackwurstking/pirgb-web/pkg/router"
)

func ServeApi(pattern string, mux *router.RegexRouter, devices pirgb.Devices) *router.RegexRouter {
	handler.Devices = devices

	mux.HandleFunc(pattern+"/events", handler.Events)

	regex, _ := regexp.Compile(pattern + "/devices(.*)")
	mux.HandleRegEx(regex, handler.NewDeviceHandler(pattern+"/devices"))

	return mux
}
