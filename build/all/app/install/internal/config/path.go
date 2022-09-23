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
	if os.Getuid() == -1 {
		panic("Windows user detected!!!!!!")
	}

	var dir string

	if os.Getenv("USER") == "root" {
		dir = "/etc/xdg" // NOTE: not using `XDG_CONFIG_DIRS` here
		logrus.Warnf("[config] Running with sudo, config directory set to \"%s\"", dir)
	} else {
		dir = os.Getenv("XDG_CONFIG_HOME")
		if dir == "" {
			home, err := os.UserHomeDir()
			if err != nil {
				logrus.Panicln(err)
			}

			dir = filepath.Join(home, ".config")
		}
	}

	return filepath.Join(dir, Vendor, Project)
}
