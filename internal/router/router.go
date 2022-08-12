package router

import (
	"github.com/go-chi/chi/v5"
)

var Mux *chi.Mux

func init() {
	Mux = chi.NewRouter()
}

// Initialize router
func Initialize() *chi.Mux {
	// TODO: do this init stuff for each route
	ui("/", Mux)
	api("/api", Mux)

	return mux
}
