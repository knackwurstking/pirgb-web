//go:build !dev

// Pirgb Webserver
package main

import (
	"flag"

	"github.com/sirupsen/logrus"
)

const (
	port = 50831
)

var (
	debug     bool
	formatter = &logrus.TextFormatter{
		DisableQuote:              true,
		DisableTimestamp:          true,
		DisableSorting:            false,
		DisableLevelTruncation:    true,
		EnvironmentOverrideColors: true,
		ForceColors:               true,
		PadLevelText:              true,
	}
)

func init() {
	// NOTE: i don't care about the `--log-level` flag for now
	flag.BoolVar(&debug, "debug", debug, "enable debug logs")
}
