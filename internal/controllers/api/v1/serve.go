package v1

import (
	"net/http"
	"regexp"

	"github.com/knackwurstking/pirgb-web/internal/controllers"
	"github.com/knackwurstking/pirgb-web/pkg/log"
	"github.com/knackwurstking/pirgb-web/pkg/middleware"
)

type DeviceHandler struct {
	ReDeviceSection *regexp.Regexp
}

func NewDeviceHandler() *DeviceHandler {
	reDeviceSection, err := regexp.Compile(`/([a-zA-z0-9 ]+)/([0-9]{1})`)
	if err != nil {
		panic(err)
	}

	return &DeviceHandler{
		ReDeviceSection: reDeviceSection,
	}
}

func (h *DeviceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	middleware.Logger(
		func(w http.ResponseWriter, r *http.Request) {
			switch path := r.URL.Path; {
			case h.ReDeviceSection.MatchString(path):
				subStrings := h.ReDeviceSection.FindStringSubmatch(path)
				log.Debug.Printf("%#v", subStrings)

				// TODO: ...

				w.WriteHeader(http.StatusOK)
			default:
				w.WriteHeader(http.StatusNotFound)
			}
		},
	)(w, r)
}

func ServeApi(pattern string, mux *controllers.RegExpHandler) *controllers.RegExpHandler {
	mux.HandleFunc(
		pattern+"/events",
		middleware.Logger(func(w http.ResponseWriter, r *http.Request) {
			// TODO: ...
		}),
	)

	regex, _ := regexp.Compile(pattern + "/devices/(.*)")
	mux.HandleRegEx(regex, NewDeviceHandler())

	return mux
}
