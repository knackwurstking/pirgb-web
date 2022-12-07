package handler

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/knackwurstking/pirgb-web/pkg/log"
)

var (
	RegexDeviceSection *regexp.Regexp
	RegexDevice        *regexp.Regexp
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
}

type DeviceHandler struct {
	RegexDeviceSection *regexp.Regexp
	RegexDevice        *regexp.Regexp
	Pattern            string
}

func NewDeviceHandler(pattern string) *DeviceHandler {
	return &DeviceHandler{
		RegexDeviceSection: RegexDeviceSection,
		RegexDevice:        RegexDevice,
		Pattern:            pattern,
	}
}

func (h *DeviceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch path := strings.Replace(r.URL.Path, h.Pattern, "", 1); {
	case path == "" || path == "/":

		switch r.Method {
		case http.MethodGet:
			// TODO: return devices data
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

		default:
			w.WriteHeader(http.StatusNotFound)
		}

	case h.RegexDeviceSection.MatchString(path):
		subStrings := h.RegexDeviceSection.FindStringSubmatch(path)
		log.Debug.Printf("%#v", subStrings)

		// TODO: ...

		w.WriteHeader(http.StatusOK)

	case h.RegexDevice.MatchString(path):
		subStrings := h.RegexDevice.FindStringSubmatch(path)
		log.Debug.Printf("%#v", subStrings)

		// TODO: ...

		w.WriteHeader(http.StatusOK)

	default:
		w.WriteHeader(http.StatusNotFound)
	}
}
