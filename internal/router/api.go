// Package router contains all the router stuff
package router

import (
	"github.com/go-chi/chi/v5"
)

func init() {
	Mux.Route("/api", func(r chi.Router) {
		// TODO: - sections => GET: "/api/sections"
		// TODO: - groups   => GET: "/api/groups"
	})

	info["GET /api/sections"] = "get all sections"
	info["GET /api/groups"] = "get all groups"
}
