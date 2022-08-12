//go:build !dev
// +build !dev

package config

import (
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

const (
	Vendor     = "knackwurstking" // Vendor name
	Project    = "pirgb-web"      // Project name
	ConfigFile = "config.yaml"    // ConfigFile name
)

func userConfigDir() string {
	dir, err := os.UserConfigDir()
	if err != nil {
		logrus.Warnf("Load config failed: %+v", err)
		return dir
	}

	return filepath.Join(dir, Vendor, Project)
}
