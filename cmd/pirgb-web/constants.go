//go:build !dev

package main

import (
	"flag"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	debug     bool
	formatter = &logrus.TextFormatter{
		DisableQuote:              true,
		DisableSorting:            true,
		DisableLevelTruncation:    true,
		EnvironmentOverrideColors: true,
		FullTimestamp:             true,
		TimestampFormat:           time.Now().Format(time.UnixDate),
	}
)

func init() {
	// NOTE: i don't care about the `--log-level` flag for now
	flag.BoolVar(&debug, "debug", debug, "enable debug logs")
}
