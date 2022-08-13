//go:build !dev

package router

import (
	"github.com/go-chi/chi/v5"
)

func init() {
	mux.Route("/", func(r chi.Router) {
		// TODO: load ui endpoints here
	})
}
