package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"

	"gitlab.com/knackwurstking/pirgb-web/internal/config"
	"gitlab.com/knackwurstking/pirgb-web/internal/router"

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(formatter)

	// parse flags here (host, port)
	flag.StringVar(&config.GlobalData.Host, "host", config.GlobalData.Host,
		"whatever ...")

	flag.IntVar(&config.GlobalData.Port, "port", config.GlobalData.Port,
		"port to bind the server to")
}

func main() {
	config.DoIt()
	flag.Parse()

	if debug {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}

	mux := router.GetMux()

	// initialize the router and server
	server := &http.Server{
		Addr: fmt.Sprintf("%s:%d",
			config.GlobalData.Host, config.GlobalData.Port),
		Handler: mux,
	}

	var wg sync.WaitGroup

	if config.GlobalData.EnableHTTP {
		wg.Add(1)
		go startServerHTTP(server, &wg)
	}

	if config.GlobalData.EnableHTTPS {
		wg.Add(1)
		go startServerHTTPS(server, &wg)
	}

	wg.Wait()
}

func startServerHTTP(server *http.Server, wg *sync.WaitGroup) {
	defer wg.Done()

	logrus.WithField("Address", server.Addr).Infof("HTTP server running ...")
	if err := server.ListenAndServe(); err != nil {
		logrus.Errorf("http: %s", err.Error())
	}
}

func startServerHTTPS(server *http.Server, wg *sync.WaitGroup) {
	defer wg.Done()

	logrus.Warnf("HTTPS server not running - work in progress")
	// TODO: need a cert first
}
