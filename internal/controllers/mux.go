package controllers

import (
	"net/http"
	"regexp"
)

type RegExpHandler struct {
	ReDeviceSection *regexp.Regexp
}

func NewRegExpHandler() *RegExpHandler {
	reDeviceSection, err := regexp.Compile(`/([a-zA-z0-9 ]+)/([0-9]{1})`)
	if err != nil {
		panic(err)
	}

	return &RegExpHandler{
		ReDeviceSection: reDeviceSection,
	}
}

func (h *RegExpHandler) Handle(pattern *regexp.Regexp, handler http.Handler) {
	// ...
}

func (h *RegExpHandler) HandleFunc(
	pattern *regexp.Regexp,
	handler func(http.ResponseWriter, *http.Request),
) {
	// ...
}

func (h *RegExpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// ...
}
