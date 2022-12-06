package handler

import "net/http"

func Events(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
