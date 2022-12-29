package router

import (
	"net/http"
	"regexp"
)

type Route struct {
	Pattern      string
	RegexPattern *regexp.Regexp // if not nil => Pattern will be ignored
	Handler      http.Handler
	Routes       []string
}

func (r *Route) IsRegex() bool {
	return r.RegexPattern != nil
}
