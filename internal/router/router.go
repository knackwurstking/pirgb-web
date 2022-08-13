package router

import (
	"fmt"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

var (
	mux = chi.NewRouter()
	info = make(map[string]string)
)

func GetMux() *chi.Mux {
	return mux
}

func PrintInfo() {
	i := "Endpoints:\n"

	for e, d := range info {
		i += fmt.Sprintf("%-30s - %s\n", e, d)
	}

	logrus.Infoln(i)
}
