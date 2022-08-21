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
	"gitlab.com/knackwurstking/pirgb-web/internal/events"
	"nhooyr.io/websocket"
)

func init() {
	Mux.Route("/api", func(r chi.Router) {
		r.Get("/events", func(w http.ResponseWriter, r *http.Request) {
			conn, err := websocket.Accept(w, r, &websocket.AcceptOptions{
				InsecureSkipVerify: true,
			})
			if err != nil {
				logrus.Errorln("[router]", err.Error())
				http.Error(w, "websocket connection failed", http.StatusInternalServerError)
				return
			}
			ctx := r.Context()
			addr := r.RemoteAddr
			events.Global.AddClient(ctx, conn, addr)

			// Connection closed handling
			defer func() {
				conn.Close(websocket.StatusAbnormalClosure, "connection aborted")
				events.Global.RemoveClientAddr(addr)
			}()

			for {
				_, _, err := conn.Read(r.Context())
				if err != nil {
					// TODO: check error type ...
					logrus.Errorf("[router] Connection read error: \"%+v\", %T", err, err)
					return
				}
			}
		})

		r.Route("/devices", func(r chi.Router) {
			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				w.Header().Add("Content-Type", "application/json")
				err := json.NewEncoder(w).Encode(config.Global.Devices)
				if err != nil {
					logrus.Warnln("[router]", err.Error())
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
