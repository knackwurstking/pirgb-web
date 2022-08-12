//go:build dev

package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

func ui(endpoint string, mux *chi.Mux) {
	logrus.Infoln("router: skip ui on dev build ...")
}
