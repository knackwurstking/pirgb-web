package router

import (
	"fmt"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/sirupsen/logrus"
)

var (
	// Mux chi router to use
	Mux         = newRouter()
	DefaultCORS = cors.Options{
		//AllowedOrigins: []string{"https://foo.com"}, // use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		//AllowOriginFunc: func(r *http.Request, origin string) bool { return true }, // allow all
		AllowedMethods: []string{"GET"},
		AllowedHeaders: []string{},
		//ExposedHeaders:   []string{"Content-Type"},
		ExposedHeaders:   []string{},
		AllowCredentials: false,
		//MaxAge:           300, // Maximum value not ignored by any of major browsers
	}
	Info = make(map[string]string)
)

// PrintInfo about router endpoints loaded
func PrintInfo() {
	i := "Endpoints:\n"

	for e, d := range Info {
		i += fmt.Sprintf("%-30s - %s\n", e, d)
	}

	logrus.Infoln(i)
}

func newRouter() *chi.Mux {
	mux := chi.NewRouter()
	mux.Use(cors.Handler(DefaultCORS))

	return mux
}
