package middleware

import (
	"net/http"

	"github.com/knackwurstking/pirgb-web/pkg/log"
)

// Authorization middleware
func Authorization(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ww := GetCachedResponse(w)

		log.Debug.Printf("PROTECTED ROUTE: %s", r.URL.Path)
		// TODO: check for basic authorization here before `ServeHTTP`

		h.ServeHTTP(ww, r)
	}
}
