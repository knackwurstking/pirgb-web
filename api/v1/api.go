package api

import (
    "net/http"
)

func NewRouter() *http.ServeMux {
    return http.NewServeMux()
}

func NewHandler(router *http.ServeMux) http.Handler {
    // ...

    return router 
}
