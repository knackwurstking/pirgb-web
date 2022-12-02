package v1

import (
	"net/http"
	"regexp"
)

type DeviceHandler struct {
	ReDeviceSection *regexp.Regexp
}

func NewDeviceHandler() *DeviceHandler {
	reDeviceSection, err := regexp.Compile(`/([a-zA-z0-9 ]+)/([0-9])`)
	if err != nil {
		panic(err)
	}

	return &DeviceHandler{
		ReDeviceSection: reDeviceSection,
	}
}

func (h *DeviceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch path := r.URL.Path; {
	case h.ReDeviceSection.MatchString(path):
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func ServeApi(pattern string, mux *http.ServeMux) *http.ServeMux {
	mux.HandleFunc(pattern+"/events", func(w http.ResponseWriter, r *http.Request) {})
	mux.Handle(pattern+"/devices", NewDeviceHandler())

	mux.HandleFunc(
		pattern+"/devices/:host/:section",
		func(w http.ResponseWriter, r *http.Request) {},
	)
	return mux
}
