package handler

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/knackwurstking/pirgb-web/internal/constants"
	"github.com/knackwurstking/pirgb-web/pkg/log"
	"github.com/knackwurstking/pirgb-web/pkg/pirgb"
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
		h.handler(w, r)

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

func (h *DeviceHandler) WriteJSON(w http.ResponseWriter, data any) {
	w.Header().Add("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(data)
	w.WriteHeader(http.StatusOK)
}

func (h *DeviceHandler) handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		h.WriteJSON(w, constants.Config.Devices)

	default:
		w.WriteHeader(http.StatusNotFound)

	}
}

func (h *DeviceHandler) handlerDevice(w http.ResponseWriter, r *http.Request) {
	host := h.RegexDevice.FindStringSubmatch(r.URL.Path)[0]

	switch r.Method {

	case http.MethodGet:
		if device := constants.Config.Devices.Get(host); device != nil {
			h.WriteJSON(w, device)
			return
		}

		w.WriteHeader(http.StatusNotFound)

	default:
		w.WriteHeader(http.StatusNotFound)

	}
}

func (h *DeviceHandler) handlerDeviceSection(w http.ResponseWriter, r *http.Request) {
	ps := h.RegexDeviceSection.FindStringSubmatch(r.URL.Path)
	host := ps[0]
	sectionID, _ := strconv.Atoi(ps[1])

	switch r.Method {

	case http.MethodGet:
		device := constants.Config.Devices.Get(host)
		if device == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		section := device.GetSection(sectionID)
		if section == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		h.WriteJSON(w, section)

	default:
		w.WriteHeader(http.StatusNotFound)

	}
}

func (h *DeviceHandler) handlerDeviceSectionPWM(w http.ResponseWriter, r *http.Request) {
	ps := h.RegexDeviceSectionPWM.FindStringSubmatch(r.URL.Path)
	host := ps[0]
	sectionID, _ := strconv.Atoi(ps[1])

	device := constants.Config.Devices.Get(host)
	if device == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	section := device.GetSection(sectionID)
	if section == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	switch r.Method {

	case http.MethodGet:
		h.WriteJSON(w, section)

	case http.MethodPost:
		var pwm pirgb.PWM

		defer r.Body.Close()
		err := json.NewDecoder(r.Body).Decode(&pwm)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = device.SetPWM(sectionID, pwm)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)

	default:
		w.WriteHeader(http.StatusNotFound)

	}
}
