package router

import (
	"github.com/go-chi/chi/v5"
)

var mux = chi.NewRouter()

func GetMux() *chi.Mux {
	return mux
}
