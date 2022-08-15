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

		r.Get("/sections", getSectionsHandler)
		r.Get("/groups", getGroupsHandler)
	})

	Info["    GET /api/sections"] = "get all sections"
	Info["    GET /api/groups"] = "get all groups"
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
