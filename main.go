package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/knackwurstking/pirgb-web/internal/constants"
	v1 "github.com/knackwurstking/pirgb-web/internal/controllers/api/v1"
	"github.com/knackwurstking/pirgb-web/internal/controllers/fileserver"
	"github.com/knackwurstking/pirgb-web/internal/events"
	"github.com/knackwurstking/pirgb-web/internal/router"
	"github.com/knackwurstking/pirgb-web/pkg/log"
)

func init() {
	// load config
	_ = constants.LoadConfig()

	flag.StringVar(&constants.Config.Host, "host", constants.Config.Host, "whatever ...")
	flag.IntVar(&constants.Config.Port, "port", constants.Config.Port, "port to bind the server to")

	flag.Parse()

	go startDeviceScan()
}

func main() {
	// initialize the router and server
	mux := router.NewRegexRouter()

	// router
	v1.ServeApi("/api/v1", mux)
	fileserver.ServeFiles(`/?`, mux)

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", constants.Config.Host, constants.Config.Port),
		Handler: mux,
	}

	log.Info.Printf("server running %s", server.Addr)
	// TODO: take from env
	if err := server.ListenAndServeTLS("cert.pem", "cert-key.pem"); err != nil {
		log.Error.Fatalf("server error: %s", err.Error())
	}
}

func startDeviceScan() {
	log.Debug.Println("start device scan...")
	constants.Config.Devices.Scan()
	events.Start()
}
