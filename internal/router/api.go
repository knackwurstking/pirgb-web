// Package router contains all the router stuff
package router

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"gitlab.com/knackwurstking/pirgb-web/internal/config"
	"gitlab.com/knackwurstking/pirgb-web/internal/events"
	"nhooyr.io/websocket"
)

func init() {
	Mux.Route("/api", func(r chi.Router) {
		r.Get("/events/{name}", func(w http.ResponseWriter, r *http.Request) {
			name := chi.URLParam(r, "name")

			switch name {
			case "change":
				conn, err := websocket.Accept(w, r, nil)
				if err != nil {
					http.Error(w, "websocket connection failed", http.StatusInternalServerError)
					return
				}
				events.Global.AddClient(r.Context(), conn, strings.Split(r.RemoteAddr, ":")[0])
			default:
				http.Error(w, "event not found", http.StatusNotFound)
				return
			}
		})

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
