// Package router contains all the router stuff
package router

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"gitlab.com/knackwurstking/pirgb-web/internal/config"
)

func init() {
	Mux.Route("/api", func(r chi.Router) {
		r.Route("/devices", func(r chi.Router) {
			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				w.Header().Add("Content-Type", "application/json")
				err := json.NewEncoder(w).Encode(config.Global.Devices)
				if err != nil {
					logrus.Warnln(err.Error())
				}
			})

			r.Route("/{host}/{section:[0-9]}", func(r chi.Router) {
				r.Use(middlewareParseDeviceData)

				r.Get("/pwm", handlerGetServerPWM)
				r.Post("/pwm", handlerPostServerPWM)
				r.Get("/ws/pwm", handlerWebsocketPWM)
				r.Get("/ws/event/change", handlerEventChange)
			})
		})
	})
}

func middlewareParseDeviceData(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		host := chi.URLParam(r, "host")

		sectionParam := chi.URLParam(r, "section") // parse this to int
		section, err := strconv.Atoi(sectionParam)
		if err != nil {
			http.Error(w, "section id not found", http.StatusNotFound)
			return
		}

		ctx = context.WithValue(ctx, "host", host)
		ctx = context.WithValue(ctx, "section", section)

		handler.ServeHTTP(w, r.WithContext(ctx))
	})
}

func handlerGetServerPWM(w http.ResponseWriter, r *http.Request) {
	// TODO: get pwm from database
}

func handlerPostServerPWM(w http.ResponseWriter, r *http.Request) {
	// TODO: set pwm to pirgb-server
}

func handlerWebsocketPWM(w http.ResponseWriter, r *http.Request) {
	// TODO: ...
}

func handlerEventChange(w http.ResponseWriter, r *http.Request) {
	// TODO: ...
}

// **** OLD: *******************************************************************************

//func getSectionsHandler(w http.ResponseWriter, r *http.Request) {
//	// sections => GET: "/api/sections"
//	w.Header().Add("Content-Type", "application/json")
//	if err := json.NewEncoder(w).Encode(config.Global.Devices); err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//}
//
//func getGroupsHandler(w http.ResponseWriter, r *http.Request) {
//	// groups   => GET: "/api/groups"
//	w.Header().Add("Content-Type", "application/json")
//	if err := json.NewEncoder(w).Encode(config.Global.Groups.ListGroups()); err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//}
//
//func postServerPWMHandler(w http.ResponseWriter, r *http.Request) {
//	host := chi.URLParam(r, "host")
//	section := chi.URLParam(r, "section")
//
//	device := config.Global.Devices.Get(host)
//	if device == nil {
//		// error: wrong {host}
//		http.Error(w, fmt.Sprintf("device not found \"%s\"", host),
//			http.StatusBadRequest)
//		return
//	}
//
//	// Build the URL
//	url := fmt.Sprintf("http://%s:%d/pwm/%s", device.Host, device.Port, section)
//	logrus.Debugf("forward to ... %s:%d", device.Host, device.Port)
//	resp, err := http.Post(url, r.Header.Get("Content-Type"), r.Body)
//	if err != nil {
//		// NOTE: check server logs for %s:%d (see url)
//		http.Error(
//			w, http.StatusText(http.StatusInternalServerError),
//			http.StatusInternalServerError,
//		)
//		return
//	}
//	defer resp.Body.Close()
//
//	utils.CopyHeaders(w.Header(), resp.Header)
//	w.WriteHeader(resp.StatusCode)
//	io.Copy(w, resp.Body)
//}
//
//func getServerPWMHandler(w http.ResponseWriter, r *http.Request) {
//	host := chi.URLParam(r, "host")
//	section := chi.URLParam(r, "section")
//
//	device := config.Global.Devices.Get(host)
//	if device == nil {
//		// error: wrong {host}
//		http.Error(w, fmt.Sprintf("device not found for \"%s\"", host),
//			http.StatusBadRequest)
//		return
//	}
//
//	// Build the URL
//	url := fmt.Sprintf("http://%s:%d/pwm/%s", device.Host, device.Port, section)
//	logrus.Debugf("forward to ... %s:%d", device.Host, device.Port)
//	resp, err := http.Get(url)
//	if err != nil {
//		// NOTE: check server logs for %s:%d (see url)
//		http.Error(w, http.StatusText(http.StatusInternalServerError),
//			http.StatusInternalServerError)
//	}
//	defer resp.Body.Close()
//
//	utils.CopyHeaders(w.Header(), resp.Header)
//	w.WriteHeader(resp.StatusCode)
//	io.Copy(w, resp.Body)
//}
