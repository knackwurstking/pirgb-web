package middleware

import (
	"net/http"
)

func Auth(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ww := GetCachedResponse(w)

		// TODO: check authorization key...
		// ...

		h.ServeHTTP(ww, r)
	}
}
