// Package router contains all the router stuff
package router

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"gitlab.com/knackwurstking/pirgb-web/internal/config"
	"gitlab.com/knackwurstking/pirgb-web/internal/events"
	"gitlab.com/knackwurstking/pirgb-web/pkg/pirgb"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"nhooyr.io/websocket"
)

func init() {
	Mux.Route("/api", func(r chi.Router) {
		r.Get("/events", func(w http.ResponseWriter, r *http.Request) {
			logrus.Debugf("[router] register event handler %s", r.RemoteAddr)

			conn, err := websocket.Accept(w, r, &websocket.AcceptOptions{
				InsecureSkipVerify: true,
			})
			if err != nil {
				logrus.Errorln("[router]", err.Error())
				http.Error(w, "websocket connection failed", http.StatusInternalServerError)
				return
			}

			addr := r.RemoteAddr

			ctx := r.Context()

			/*
				pingCtx, pingCancel := context.WithTimeout(ctx, time.Duration(time.Second*3))
				defer pingCancel()
				go func() {
					for {
						time.Sleep(5)
						err := conn.Ping(pingCtx)
						if err != nil {
							logrus.Warnf("[router] ping failed: %s", err.Error())
							return
						}
					}
				}()
			*/

			events.Global.AddClient(ctx, conn, addr)
			defer func() {
				conn.Close(websocket.StatusAbnormalClosure, "connection aborted")
				events.Global.RemoveClientAddr(addr)
			}()

			for {
				msgType, msg, err := conn.Read(r.Context())
				if err != nil {
					logrus.Errorf("[router] Connection read error: \"%+v\", %T", err, err)
					return
				}

				go func() {
					err = conn.Write(r.Context(), msgType, msg)
					if err != nil {
						logrus.Warnf("[router] Send echo failed: %s", err.Error())
					}
				}()
			}
		})

		r.Route("/devices", func(r chi.Router) {
			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				err := json.NewEncoder(w).Encode(config.Global.Devices)
				if err != nil {
					logrus.Warnln("[router]", err.Error())
					w.WriteHeader(http.StatusNoContent)
					return
				}
				w.Header().Add("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
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
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		ctx = context.WithValue(ctx, "host", host)
		ctx = context.WithValue(ctx, "section", section)

		handler.ServeHTTP(w, r.WithContext(ctx))
	})
}

func handlerGetServerPWM(w http.ResponseWriter, r *http.Request) {
	host := r.Context().Value("host").(string)
	sectionID := r.Context().Value("section").(int)

	device := config.Global.Devices.Get(host)
	if device == nil {
		http.Error(w, "device not found", http.StatusNotFound)
		return
	}

	section := device.GetSection(sectionID)
	if section == nil {
		http.Error(w, "section not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(section)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func handlerPostServerPWM(w http.ResponseWriter, r *http.Request) {
	host := r.Context().Value("host").(string)
	sectionID := r.Context().Value("section").(int)

	//  get request
	var pwmData pirgb.PWM
	err := json.NewDecoder(r.Body).Decode(&pwmData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	device := config.Global.Devices.Get(host)
	if device == nil {
		http.Error(w, "device not found", http.StatusNotFound)
		return
	}

	err = device.SetPWM(sectionID, pwmData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
