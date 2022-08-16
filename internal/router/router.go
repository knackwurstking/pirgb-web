package router

import (
	"fmt"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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
	Info = make(Endpoints, 0)
)

type Endpoints []EndpointInfo

func (Endpoints) Print() {
	i := "Endpoints:\n"

	for _, info := range Info {
		i += fmt.Sprintf("    %7s %-23s %s\n", info.Method, info.Endpoint, info.Description)
	}

	logrus.Infoln(i)
}

type EndpointInfo struct {
	Method      string
	Endpoint    string
	Description string
}

func NewEndpointInfo(method, endpoint, desc string) EndpointInfo {
	return EndpointInfo{
		Method:      method,
		Endpoint:    endpoint,
		Description: desc,
	}
}

func newRouter() *chi.Mux {
	mux := chi.NewRouter()
	mux.Use(middleware.Logger)
	mux.Use(cors.Handler(DefaultCORS))

	return mux
}
