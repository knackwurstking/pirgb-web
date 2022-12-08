package handler

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/knackwurstking/pirgb-web/pkg/log"
)

var (
	RegexDevice           *regexp.Regexp
	RegexDeviceSection    *regexp.Regexp
	RegexDeviceSectionPWM *regexp.Regexp
)

func init() {
	var err error

	RegexDevice, err = regexp.Compile(`/([a-zA-z0-9 ]+)/?`)
	if err != nil {
		log.Error.Panicln(err)
	}

	RegexDeviceSection, err = regexp.Compile(`/([a-zA-z0-9 ]+)/([0-9]{1})/?`)
	if err != nil {
		log.Error.Panicln(err)
	}

	RegexDeviceSectionPWM, err = regexp.Compile(`/([a-zA-z0-9 ]+)/([0-9]{1})/pwm/?`)
	if err != nil {
		log.Error.Panicln(err)
	}
}

type DeviceHandler struct {
	RegexDevice           *regexp.Regexp
	RegexDeviceSection    *regexp.Regexp
	RegexDeviceSectionPWM *regexp.Regexp
	Pattern               string
}

func NewDeviceHandler(pattern string) *DeviceHandler {
	return &DeviceHandler{
		RegexDevice:           RegexDevice,
		RegexDeviceSection:    RegexDeviceSection,
		RegexDeviceSectionPWM: RegexDeviceSection,
		Pattern:               pattern,
	}
}

func (h *DeviceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch path := strings.Replace(r.URL.Path, h.Pattern, "", 1); {

	// "/api/v1/devices/"
	case path == "" || path == "/":

		switch r.Method {
		case http.MethodGet:
			h.handler(w, r)
			w.WriteHeader(http.StatusOK)

		default:
			w.WriteHeader(http.StatusNotFound)
		}

	// "/api/v1/devices/:device"
	case h.RegexDevice.MatchString(path):
		subStrings := h.RegexDevice.FindStringSubmatch(path)
		log.Debug.Printf("%#v", subStrings)

		h.handlerDevice(w, r)

	// "/api/v1/devices/:device/:section"
	case h.RegexDeviceSection.MatchString(path):
		subStrings := h.RegexDeviceSection.FindStringSubmatch(path)
		log.Debug.Printf("%#v", subStrings)

		h.handlerDeviceSection(w, r)

	// "/api/v1/devices/:device/:section"
	case h.RegexDeviceSectionPWM.MatchString(path):
		subStrings := h.RegexDeviceSectionPWM.FindStringSubmatch(path)
		log.Debug.Printf("%#v", subStrings)

		h.handlerDeviceSectionPWM(w, r)

	default:
		w.WriteHeader(http.StatusNotFound)

	}
}

func (h *DeviceHandler) handler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:
		w.Header().Add("Content-Type", "application/json")

		// TODO: return all devices data
		// ...

		w.WriteHeader(http.StatusOK)

	default:
		w.WriteHeader(http.StatusNotFound)

	}
}

func (h *DeviceHandler) handlerDevice(w http.ResponseWriter, r *http.Request) {
	// TODO: ...
}

func (h *DeviceHandler) handlerDeviceSection(w http.ResponseWriter, r *http.Request) {
	// TODO: ...
}

func (h *DeviceHandler) handlerDeviceSectionPWM(w http.ResponseWriter, r *http.Request) {
	// TODO: ...
}
