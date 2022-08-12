//go:build !dev

package router

import "github.com/go-chi/chi/v5"

func ui(endpoint string, mux *chi.Mux) {
	// TODO: embed ui files
}
