package router

import (
	"fmt"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

var (
	Mux = chi.NewRouter()
	info = make(map[string]string)
)

func PrintInfo() {
	i := "Endpoints:\n"

	for e, d := range info {
		i += fmt.Sprintf("%-30s - %s\n", e, d)
	}

	logrus.Infoln(i)
}
