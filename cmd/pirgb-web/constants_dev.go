//go:build dev

package main

import "github.com/sirupsen/logrus"

var (
	debug     = true
	formatter = &logrus.TextFormatter{
		DisableQuote:              true,
		DisableTimestamp:          true,
		DisableSorting:            true,
		DisableLevelTruncation:    true,
		EnvironmentOverrideColors: true,
	}
)
