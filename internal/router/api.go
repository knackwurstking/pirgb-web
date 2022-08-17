// Package router contains all the router stuff
package router

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"gitlab.com/knackwurstking/pirgb-web/internal/config"
	"gitlab.com/knackwurstking/pirgb-web/internal/utils"
)

func init() {
	Mux.Route("/api", func(r chi.Router) {

		r.Route("/devices", func(r chi.Router) {
			r.Get("/", getSectionsHandler)
			r.Post("/{host}/pwm/{section:[0-9]}", postServerPWMHandler)
			r.Get("/{host}/pwm/{section:[0-9]}", getServerPWMHandler)
		})

		r.Get("/groups", getGroupsHandler)
	})

	Info = append(Info, NewEndpointInfo("GET", "/api/devices", "get devices"))
	Info = append(Info, NewEndpointInfo("GET", "/api/groups", "get available groups"))

	Info = append(Info, NewEndpointInfo("POST", "/api/devices/{host}/pwm/{section:[0-9]}",
		"forwards the request to the pirgb-server (device)"))
	Info = append(Info, NewEndpointInfo("GET", "/api/devices/{host}/pwm/{section:[0-9]}",
		"forwards the request to the pirgb-server (device)"))
}

func getSectionsHandler(w http.ResponseWriter, r *http.Request) {
	// sections => GET: "/api/sections"
	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(config.Global.Devices); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getGroupsHandler(w http.ResponseWriter, r *http.Request) {
	// groups   => GET: "/api/groups"
	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(config.Global.Groups.ListGroups()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func postServerPWMHandler(w http.ResponseWriter, r *http.Request) {
	host := chi.URLParam(r, "host")
	section := chi.URLParam(r, "section")

	device := config.Global.Devices.Get(host)
	if device == nil {
		// error: wrong {host}
		http.Error(w, fmt.Sprintf("device not found \"%s\"", host),
			http.StatusBadRequest)
		return
	}

	// Build the URL
	url := fmt.Sprintf("http://%s:%d/pwm/%s", device.Host, device.Port, section)
	logrus.Debugf("forward to ... %s:%d", device.Host, device.Port)
	resp, err := http.Post(url, r.Header.Get("Content-Type"), r.Body)
	if err != nil {
		// NOTE: check server logs for %s:%d (see url)
		http.Error(
			w, http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
		return
	}
	defer resp.Body.Close()

	utils.CopyHeaders(w.Header(), resp.Header)
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

func getServerPWMHandler(w http.ResponseWriter, r *http.Request) {
	host := chi.URLParam(r, "host")
	section := chi.URLParam(r, "section")

	device := config.Global.Devices.Get(host)
	if device == nil {
		// error: wrong {host}
		http.Error(w, fmt.Sprintf("device not found for \"%s\"", host),
			http.StatusBadRequest)
		return
	}

	// Build the URL
	url := fmt.Sprintf("http://%s:%d/pwm/%s", device.Host, device.Port, section)
	logrus.Debugf("forward to ... %s:%d", device.Host, device.Port)
	resp, err := http.Get(url)
	if err != nil {
		// NOTE: check server logs for %s:%d (see url)
		http.Error(w, http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
	}
	defer resp.Body.Close()

	utils.CopyHeaders(w.Header(), resp.Header)
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}
