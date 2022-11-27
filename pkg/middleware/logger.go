package middleware

import (
	"net/http"
	"time"

	"simple-pwa-example/pkg/log"
)

// Logger - endpoint logger
func Logger(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ww := GetCachedResponse(w)

		defer func(ww *ResponseWriter, start time.Time) {
			log.Info.Printf("%-6s %s from %s - %d %s - %s",
				r.Method, r.URL.Path, r.RemoteAddr,
				ww.Status, http.StatusText(ww.Status),
				time.Since(start).String())
		}(ww, time.Now())

		h.ServeHTTP(ww, r)
	}
}
