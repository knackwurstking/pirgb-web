package controllers

import (
	"net/http"
	"regexp"
)

type Route struct {
	Pattern       string
	RegExpPattern *regexp.Regexp // if not nil => Pattern will be ignored
	Handler       http.Handler
}

func (r *Route) IsRegExp() bool {
	return r.RegExpPattern != nil
}

type RegExpHandler struct {
	Routes []*Route
}

func NewRegExpHandler() *RegExpHandler {
	return &RegExpHandler{
		Routes: make([]*Route, 0),
	}
}

func (h *RegExpHandler) HandleRegEx(pattern *regexp.Regexp, handler http.Handler) {
	h.Routes = append(h.Routes, &Route{RegExpPattern: pattern, Handler: handler})
}

func (h *RegExpHandler) Handle(pattern string, handler http.Handler) {
	h.Routes = append(h.Routes, &Route{Pattern: pattern, Handler: handler})
}

func (h *RegExpHandler) HandleRegExFunc(
	pattern *regexp.Regexp,
	handler func(http.ResponseWriter, *http.Request),
) {
	h.Routes = append(h.Routes, &Route{RegExpPattern: pattern, Handler: http.HandlerFunc(handler)})
}

func (h *RegExpHandler) HandleFunc(
	pattern string,
	handler func(http.ResponseWriter, *http.Request),
) {
	h.Routes = append(h.Routes, &Route{Pattern: pattern, Handler: http.HandlerFunc(handler)})
}

func (h *RegExpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO: iter routes - match pattern run handler return or if nothing found `http.NotFound(w, r)`
}
