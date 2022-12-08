package router

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/knackwurstking/pirgb-web/pkg/middleware"
)

type Route struct {
	Pattern      string
	RegexPattern *regexp.Regexp // if not nil => Pattern will be ignored
	Handler      http.Handler
}

func (r *Route) IsRegex() bool {
	return r.RegexPattern != nil
}

type RegexRouter struct {
	Routes []*Route
}

func NewRegexRouter() *RegexRouter {
	return &RegexRouter{
		Routes: make([]*Route, 0),
	}
}

func (h *RegexRouter) HandleRegEx(
	pattern *regexp.Regexp,
	handler http.Handler,
) {
	h.Routes = append(
		h.Routes,
		&Route{
			RegexPattern: pattern,
			Handler:      handler,
		},
	)
}

func (h *RegexRouter) Handle(pattern string, handler http.Handler) {
	h.Routes = append(h.Routes, &Route{Pattern: pattern, Handler: handler})
}

func (h *RegexRouter) HandleRegExFunc(
	pattern *regexp.Regexp,
	handler func(http.ResponseWriter, *http.Request),
) {
	h.Routes = append(h.Routes, &Route{RegexPattern: pattern, Handler: http.HandlerFunc(handler)})
}

func (h *RegexRouter) HandleFunc(
	pattern string,
	handler func(http.ResponseWriter, *http.Request),
) {
	h.Routes = append(
		h.Routes,
		&Route{
			Pattern: pattern,
			Handler: http.HandlerFunc(handler),
		},
	)
}

func (h *RegexRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	middleware.Logger(
		middleware.CORS(
			func(w http.ResponseWriter, r *http.Request) {
				for _, route := range h.Routes {
					if route.IsRegex() {
						if route.RegexPattern.MatchString(r.URL.Path) {
							route.Handler.ServeHTTP(w, r)
							return
						}
					} else if strings.TrimRight(route.Pattern, "/") == strings.TrimRight(r.URL.Path, "/") {
						route.Handler.ServeHTTP(w, r)
						return
					}
				}
			},
		),
	)(w, r)
}
