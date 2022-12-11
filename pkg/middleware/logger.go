package middleware

import (
	"net/http"
	"time"

	"github.com/knackwurstking/pirgb-web/pkg/log"
)

// Logger - endpoint logger
func Logger(h http.HandlerFunc) http.HandlerFunc {
	logger := func(w *ResponseWriter, r *http.Request, start time.Time) {
		log.Info.Printf("%-6s %s from %s - %d %s - %s",
			r.Method, r.URL.Path, r.RemoteAddr,
			w.Status, http.StatusText(w.Status),
			time.Since(start).String())
	}

	return func(w http.ResponseWriter, r *http.Request) {
		ww := GetCachedResponse(w)
		defer logger(ww, r, time.Now())

		h.ServeHTTP(ww, r)
	}
}
