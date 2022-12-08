package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/knackwurstking/pirgb-web/internal/constants"
	v1 "github.com/knackwurstking/pirgb-web/internal/controllers/api/v1"
	"github.com/knackwurstking/pirgb-web/internal/controllers/fileserver"
	"github.com/knackwurstking/pirgb-web/internal/events"

	"github.com/knackwurstking/pirgb-web/pkg/log"
	"github.com/knackwurstking/pirgb-web/pkg/router"

	aliceConfig "github.com/knackwurstking/alice/pkg/config"
)

var c *constants.Config

func init() {
	// load config
	c, _ = constants.LoadConfig()

	flag.StringVar(&c.Host, "host", c.Host, "whatever ...")
	flag.IntVar(&c.Port, "port", c.Port, "port to bind the server to")

	flag.Parse()

	// Add server to the alice config.json
	aliceConfig.SetServer(&aliceConfig.Server{Name: constants.ApplicationName, Port: c.Port})
	go startDeviceScan()
	go startEventHandler()
}

func main() {
	// start data stuff (and event handlers)
	events.Initialize(c)

	// initialize the router and server
	mux := router.NewRegexRouter()

	// router
	v1.ServeApi("/api/v1", mux)
	fileserver.ServeFiles("/", mux)

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", c.Host, c.Port),
		Handler: mux,
	}

	log.Info.Printf("server running %s", server.Addr)
	if err := server.ListenAndServeTLS("cert.crt", "cert.key"); err != nil {
		log.Error.Fatalf("server error: %s", err.Error())
	}
}

func startDeviceScan() {
	// TODO: scan devices for section data
}

func startEventHandler() {
	// TODO: start event handler for devices (sections)
}
