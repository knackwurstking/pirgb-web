package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

var (
	// Mux chi router to use
	Mux         = newRouter()
	DefaultCORS = cors.Options{
		//AllowedOrigins: []string{"https://foo.com"}, // use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		//AllowOriginFunc: func(r *http.Request, origin string) bool { return true }, // allow all
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type"},
		ExposedHeaders: []string{"Content-Type"},
		//ExposedHeaders:   []string{},
		AllowCredentials: false,
		//MaxAge:           300, // Maximum value not ignored by any of major browsers
	}
)

func newRouter() *chi.Mux {
	mux := chi.NewRouter()
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	mux.Use(cors.Handler(DefaultCORS))

	return mux
}
