// Package middleware provides middleware for the net/http module.
package middleware

import (
	"bufio"
	"errors"
	"net"
	"net/http"
)

var (
	// Cache stores all created ResponseWriter objects
	Cache []*ResponseWriter
)

// ResponseWriter for keep track of the status code.
type ResponseWriter struct {
	http.ResponseWriter

	Status int
}

// NewResponseWriter returns a new, or cached (if available), ResponseWriter. (will be added to Cache)
func NewResponseWriter(responseWriter http.ResponseWriter) *ResponseWriter {
	w := &ResponseWriter{
		ResponseWriter: responseWriter,
	}

	Cache = append(Cache, w)
	return w
}

func (w *ResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	h, ok := w.ResponseWriter.(http.Hijacker)
	if !ok {
		return nil, nil, errors.New("hijack not supported")
	}
	return h.Hijack()
}

// WriteHeader overwrites the default WriteHeader method to keep track of the status code
func (w *ResponseWriter) WriteHeader(statusCode int) {
	w.Status = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

// GetCachedResponse returns a cached ResponseWriter or creates a new one
func GetCachedResponse(w http.ResponseWriter) *ResponseWriter {
	for _, responseWriter := range Cache {
		if responseWriter == w {
			return responseWriter
		}
	}

	return NewResponseWriter(w)
}
