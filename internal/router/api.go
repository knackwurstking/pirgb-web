// Package router contains all the router stuff
package router

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"gitlab.com/knackwurstking/pirgb-web/internal/config"
)

func init() {
	Mux.Route("/api", func(r chi.Router) {
		r.Use(cors.Handler(cors.Options{
			//AllowedOrigins: []string{"https://foo.com"}, // use this to allow specific origin hosts
			AllowedOrigins: []string{"https://*", "http://*"},
			//AllowOriginFunc: func(r *http.Request, origin string) bool { return true }, // allow all
			AllowedMethods: []string{"GET"},
			AllowedHeaders: []string{},
			//ExposedHeaders:   []string{"Content-Type"},
			ExposedHeaders:   []string{},
			AllowCredentials: false,
			//MaxAge:           300, // Maximum value not ignored by any of major browsers
		}))

		r.Get("/sections", getSectionsHandler)
		r.Get("/groups", getGroupsHandler)
	})

	info["    GET /api/sections"] = "get all sections"
	info["    GET /api/groups"] = "get all groups"
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
	if err := json.NewEncoder(w).Encode(config.Global.Groups); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
