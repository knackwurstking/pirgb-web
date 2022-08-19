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
	configFile = "config.yaml"    // ConfigFile name
)

func GetConfigPath() string {
	// TODO: handle user "root"
	dir, err := os.UserConfigDir()
	if err != nil {
		logrus.Warnf("Load config failed: %s", err.Error())
		return dir
	}

	return filepath.Join(dir, Vendor, Project)
}
