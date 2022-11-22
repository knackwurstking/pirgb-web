package main

import (
	"flag"
	"fmt"
	"net/http"

	"gitlab.com/knackwurstking/pirgb-web/internal/config"
	"gitlab.com/knackwurstking/pirgb-web/internal/constants"
	"gitlab.com/knackwurstking/pirgb-web/internal/events"
	"gitlab.com/knackwurstking/pirgb-web/internal/log"
	"gitlab.com/knackwurstking/pirgb-web/internal/router"
)

func init() {
	flag.StringVar(&constants.Host, "host", constants.Host, "whatever ...")
	flag.IntVar(&constants.Port, "port", constants.Port, "port to bind the server to")

	flag.Parse()
}

func main() {
	// start data stuff (and event handlers)
	events.Initialize()

	// initialize the router and server
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", config.Global.Host, config.Global.Port),
		Handler: router.Mux,
	}

	log.Info.Printf("server running %s", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Error.Fatalf("server error: %s", err.Error())
	}
}
