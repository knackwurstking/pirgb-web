package handler

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/knackwurstking/pirgb-web/internal/constants"
	"github.com/knackwurstking/pirgb-web/internal/middleware"
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

	RegexDevice, err = regexp.Compile(`.*/([a-zA-z0-9-_]{1,})/?$`)
	if err != nil {
		log.Error.Panicln(err.Error())
	}

	RegexDeviceSection, err = regexp.Compile(`.*/([a-zA-z0-9-_]{1,})/([0-9]{1})/?$`)
	if err != nil {
		log.Error.Panicln(err.Error())
	}

	RegexDeviceSectionPWM, err = regexp.Compile(`.*/([a-zA-z0-9-_]{1,})/([0-9]{1})/pwm/?$`)
	if err != nil {
		log.Error.Panicln(err.Error())
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
		RegexDeviceSectionPWM: RegexDeviceSectionPWM,
		Pattern:               pattern,
	}
}

func (h *DeviceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch path := strings.Replace(r.URL.Path, h.Pattern, "", 1); {

	// "/api/v1/devices/"
	case path == "" || path == "/":
		middleware.Authorization(h.handler)(w, r)

	// "/api/v1/devices/:device/:section/pwm"
	case h.RegexDeviceSectionPWM.MatchString(path):
		middleware.Authorization(h.handlerDeviceSectionPWM)(w, r)

	// "/api/v1/devices/:device/:section"
	case h.RegexDeviceSection.MatchString(path):
		middleware.Authorization(h.handlerDeviceSection)(w, r)

	// "/api/v1/devices/:device"
	case h.RegexDevice.MatchString(path):
		middleware.Authorization(h.handlerDevice)(w, r)

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
	host := h.RegexDevice.FindStringSubmatch(r.URL.Path)[1]

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
	host := ps[1]
	sectionID, _ := strconv.Atoi(ps[2])

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
	host := ps[1]
	sectionID, _ := strconv.Atoi(ps[2])

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
