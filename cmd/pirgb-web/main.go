package main

import (
	"flag"
	"fmt"
	"net/http"

	"gitlab.com/knackwurstking/pirgb-web/internal/config"
	"gitlab.com/knackwurstking/pirgb-web/internal/database"
	"gitlab.com/knackwurstking/pirgb-web/internal/router"

	"github.com/sirupsen/logrus"
)

var (
	infoFlag bool
)

func init() {
	logrus.SetFormatter(formatter)
}

func main() {
	// parse flags here (host, port)
	flag.StringVar(&config.Global.Host, "host", config.Global.Host,
		"whatever ...")

	flag.IntVar(&config.Global.Port, "port", config.Global.Port,
		"port to bind the server to")

	flag.Parse()

	if debug {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.Debugf("config path: \"%s\"", config.Global.Path)
		logrus.Debugf("%+v", config.Global)
		for _, device := range config.Global.Devices {
			logrus.Debugf("%+v", device)
		}
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}

	// start data stuff (and event handlers)
	database.InitializeGlobal()

	// initialize the router and server
	server := &http.Server{
		Addr: fmt.Sprintf("%s:%d",
			config.Global.Host, config.Global.Port),
		Handler: router.Mux,
	}

	logrus.WithField("Address", server.Addr).Infof("Server running ...")
	if err := server.ListenAndServe(); err != nil {
		logrus.Fatalln(err.Error())
	}
}
