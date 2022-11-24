package api

import "net/http"

type ResponseWriter struct {
    http.ResponseWriter

    Status int
}

func NewResponseWriter(status int, w http.ResponseWriter) *ResponseWriter {
    return &ResponseWriter{
        ResponseWriter: w,
        Status: status,
    }
}

func (w *ResponseWriter) WriteHeader(statusCode int) {
    w.Status = statusCode
}
