// TODO: outsource this to pkg/router
package controllers

import (
	"net/http"
	"regexp"
	"strings"
)

type Route struct {
	Pattern      string
	RegexPattern *regexp.Regexp // if not nil => Pattern will be ignored
	Handler      http.Handler
}

func (r *Route) IsRegex() bool {
	return r.RegexPattern != nil
}

type RegexHandler struct {
	Routes     []*Route
	Middleware []http.Handler
}

func NewRegexHandler() *RegexHandler {
	return &RegexHandler{
		Routes: make([]*Route, 0),
	}
}

func (h *RegexHandler) HandleRegEx(pattern *regexp.Regexp, handler http.Handler) {
	h.Routes = append(h.Routes, &Route{RegexPattern: pattern, Handler: handler})
}

func (h *RegexHandler) Handle(pattern string, handler http.Handler) {
	h.Routes = append(h.Routes, &Route{Pattern: pattern, Handler: handler})
}

func (h *RegexHandler) HandleRegExFunc(
	pattern *regexp.Regexp,
	handler func(http.ResponseWriter, *http.Request),
) {
	h.Routes = append(h.Routes, &Route{RegexPattern: pattern, Handler: http.HandlerFunc(handler)})
}

func (h *RegexHandler) HandleFunc(
	pattern string,
	handler func(http.ResponseWriter, *http.Request),
) {
	h.Routes = append(h.Routes, &Route{Pattern: pattern, Handler: http.HandlerFunc(handler)})
}

func (h *RegexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, route := range h.Routes {
		if route.IsRegex() {
			// TODO: run regexp check on path...
			// ...
		} else if strings.TrimRight(route.Pattern, "/") == strings.TrimRight(r.URL.Path, "/") { // TODO: Ignore ending "/"
			route.Handler.ServeHTTP(w, r)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}
