//go:build dev

package main

import "github.com/sirupsen/logrus"

const (
	port = 50831
)

var (
	debug     = true
	formatter = &logrus.TextFormatter{
		DisableQuote:              true,
		DisableTimestamp:          true,
		DisableSorting:            true,
		DisableLevelTruncation:    true,
		EnvironmentOverrideColors: true,
		PadLevelText:              true,
	}
)
