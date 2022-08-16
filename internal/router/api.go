// Package router contains all the router stuff
package router

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"gitlab.com/knackwurstking/pirgb-web/internal/config"
)

func init() {
	Mux.Route("/api", func(r chi.Router) {

		r.Get("/devices", getSectionsHandler)
		r.Get("/groups", getGroupsHandler)
		// TODO: Add device control stuff (set and get pwm, just redirect stuff to the servers?)
	})

	Info = append(Info, NewEndpointInfo("GET", "/api/devices", "get devices"))
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
