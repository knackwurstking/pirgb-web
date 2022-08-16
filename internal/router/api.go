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
)

func init() {
	Mux.Route("/api", func(r chi.Router) {

		r.Route("/devices", func(r chi.Router) {
			r.Get("/", getSectionsHandler)

			// Redirect request to pirgb-server (device)
			r.Post("/{host}/pwm/{section:[0-9]}", func(w http.ResponseWriter, r *http.Request) {
				device := config.Global.Devices.Get(chi.URLParam(r, "host"))
				if device == nil {
					http.Error(w, "device not found", http.StatusBadRequest)
					return
				}

				url := fmt.Sprintf(
					"http://%s:%d/pwm/%s",
					device.Host, device.Port,
					chi.URLParam(r, "section"),
				)

				logrus.Debugf("redirect to ... %s:%d", device.Host, device.Port)

				resp, err := http.Post(url, r.Header.Get("Content-Type"), r.Body)
				if err != nil {
					http.Error(
						w, http.StatusText(http.StatusInternalServerError),
						http.StatusInternalServerError,
					)
					return
				}
				defer resp.Body.Close()
				CopyHeader(w.Header(), resp.Header)
				w.WriteHeader(resp.StatusCode)
				io.Copy(w, resp.Body)
			})
		})

		r.Get("/groups", getGroupsHandler)
	})

	Info = append(Info, NewEndpointInfo("GET", "/api/devices", "get devices"))
	Info = append(Info, NewEndpointInfo("POST", "/api/devices/{host}/pwm/{section:[0-9]}",
		"redirects the request to the pirgb-server (device)"))
	Info = append(Info, NewEndpointInfo("GET", "/api/groups", "get available groups"))
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
