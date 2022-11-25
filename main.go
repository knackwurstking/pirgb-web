package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/knackwurstking/pirgb-web/internal/constants"
	"github.com/knackwurstking/pirgb-web/internal/events"
	"github.com/knackwurstking/pirgb-web/internal/log"
	"github.com/knackwurstking/pirgb-web/api/v1"

	aliceConfig "github.com/knackwurstking/alice/pkg/config"
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

	// Add server to the alice config.json
	aliceConfig.SetServer(&aliceConfig.Server{Name: constants.ApplicationName, Port: c.Port})
}

func main() {
	// start data stuff (and event handlers)
	events.Initialize(c)

	// initialize the router and server
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", c.Host, c.Port),
		Handler: api.NewHandler(
			api.NewRouter(),
		),
	}

	log.Info.Printf("server running %s", server.Addr)
	if err := server.ListenAndServeTLS("cert.crt", "cert.key"); err != nil {
		log.Error.Fatalf("server error: %s", err.Error())
	}
}
