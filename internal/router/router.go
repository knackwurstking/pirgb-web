package router

import (
	"github.com/go-chi/chi/v5"
)

type RouterConfig struct {
	Ui  string // Ui endpoint
	Api string // Api endpoint
}

// Initialize router
func Initialize(config RouterConfig) *chi.Mux {
	mux := chi.NewRouter()

	if config.Ui != "" {
		// server public (frontend) folder
		ui(config.Ui, mux)
	}

	if config.Api != "" {
		// api routes
		api(config.Api, mux)
	}

	return mux
}
