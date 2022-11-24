package api

import (
	"net/http"

	"gitlab.com/knackwurstking/pirgb-web/internal/log"
)

var (
    Endpoint = "/api/v1"
)

func NewRouter() *http.ServeMux {
    return http.NewServeMux()
}

func NewHandler(router *http.ServeMux) http.Handler {
    router.HandleFunc(Endpoint, Handler)

    return router
}

func Handler(_w http.ResponseWriter, r *http.Request) {
    w := NewResponseWriter(http.StatusNotFound, _w)

    if auth := Authorization(w, r); auth != nil {
        switch r.URL.Path[len(Endpoint):] {
        case "", "/":
            // TODO: deliver frontend (ui) - deliver test frontend for telegram auth
            w.WriteHeader(http.StatusNoContent)
        }
    }

    w.ResponseWriter.WriteHeader(w.Status)
    log.Info.Printf("%-6s %s from %s - %d %s",
        r.Method, r.URL.Path, r.RemoteAddr, w.Status,
        http.StatusText(w.Status))
}
