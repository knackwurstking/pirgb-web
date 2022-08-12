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

func main() {

	var debug bool

	// NOTE: i don't care about the `--log-level` flag for now
	flag.BoolVar(&debug, "debug", debug, "enable debug logs")

	// parse flags here (host, port)
	flag.StringVar(&config.GlobalData.Host, "host", config.GlobalData.Host,
		"whatever ...")

	flag.IntVar(&config.GlobalData.Port, "port", config.GlobalData.Port,
		"port to bind the server to")

	flag.Parse()

	if debug {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}

	// initialize the router and server
	server := &http.Server{
		Addr: fmt.Sprintf("%s:%d",
			config.GlobalData.Host, config.GlobalData.Port),
		Handler: router.Initialize(),
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

	if err := server.ListenAndServe(); err != nil {
		logrus.Errorf("http: %s", err.Error())
	}
}

func startServerHTTPS(server *http.Server, wg *sync.WaitGroup) {
	defer wg.Done()

	// TODO: need a cert first
}
