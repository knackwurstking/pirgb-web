package middleware

import (
	"net/http"
)

// Logger - endpoint logger
func CORS(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ww := GetCachedResponse(w)

		// TODO: cors handler...
		// ...

		h.ServeHTTP(ww, r)
	}
}
