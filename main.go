package main

import (
	"flag"
	"fmt"
	"net/http"

	"gitlab.com/knackwurstking/pirgb-web/internal/constants"
	"gitlab.com/knackwurstking/pirgb-web/internal/events"
	"gitlab.com/knackwurstking/pirgb-web/internal/log"
	"gitlab.com/knackwurstking/pirgb-web/api/v1"
)

var (
	c *constants.Config
)

func init() {
	// load config
	c, _ = constants.LoadConfig()

	flag.StringVar(&c.Host, "host", c.Host, "whatever ...")
	flag.IntVar(&c.Port, "port", c.Port, "port to bind the server to")

	flag.Parse()
}

func main() {
	// start data stuff (and event handlers)
	events.Initialize(c)

	// initialize the router and server
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", c.Host, c.Port),
		Handler: api.Initialize(
			api.NewRouter(),
		),
	}

	log.Info.Printf("server running %s", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Error.Fatalf("server error: %s", err.Error())
	}
}
