//go:build !dev
// +build !dev

package constants

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/knackwurstking/pirgb-web/pkg/log"
)

const (
	PIRGB_SERVER_PORT = 50826
)

var (
	ApplicationName = "pirgb-web"
	VendorName      = "knackwurstking"
	Config          *config
)

func LoadConfig() error {
	Config = NewDefaultConfig()
	configPath, _ := os.UserConfigDir()

	// read "config.json"
	data, err := os.ReadFile(filepath.Join(
		configPath, VendorName, ApplicationName, "config.json"))
	if err != nil {
		log.Warn.Printf("read config failed: %s", err.Error())
		return err
	}

	// parse json
	err = json.Unmarshal(data, &Config)
	if err != nil {
		log.Warn.Printf("parse json config failed: %s", err.Error())
		return err
	}

	return err
}
