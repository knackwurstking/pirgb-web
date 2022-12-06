package handler

import (
	"net/http"

	"github.com/knackwurstking/pirgb-web/pkg/log"
)

func Events(w http.ResponseWriter, r *http.Request) {
	log.Debug.Println("/events...")
	w.WriteHeader(http.StatusOK)
}
