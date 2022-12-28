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

	aliceConfig "github.com/knackwurstking/alice/pkg/config"
)

func init() {
	// load config
	constants.LoadConfig()

	flag.StringVar(&constants.Config.Host, "host", constants.Config.Host, "whatever ...")
	flag.IntVar(&constants.Config.Port, "port", constants.Config.Port, "port to bind the server to")

	flag.Parse()

	// Add server to the alice config.json
	aliceConfig.SetServer(
		&aliceConfig.Server{Name: constants.ApplicationName, Port: constants.Config.Port},
	)
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
	if err := server.ListenAndServeTLS("cert.crt", "cert.key"); err != nil {
		log.Error.Fatalf("server error: %s", err.Error())
	}
}

func startDeviceScan() {
	log.Debug.Println("start device scan...")
	constants.Config.Devices.Scan()
	events.Start()
}
