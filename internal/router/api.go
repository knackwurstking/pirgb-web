package router

import (
	"github.com/go-chi/chi/v5"
)

func init() {
	mux.Route("/api", func(r chi.Router) {
		// TODO: - sections => GET: "/api/sections"
		// TODO: - groups   => GET: "/api/groups"
	})
}
