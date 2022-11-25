//go:build !dev
// +build !dev

package constants

import (
	"encoding/json"
	"os"
	"path/filepath"

	"gitlab.com/knackwurstking/pirgb-web/internal/log"
)

const (
	PIRGB_SERVER_PORT = 50826
	APPLICATION_NAME  = "pirgb-web"
	VENDOR_NAME       = "knackwurstking"
)

func LoadConfig() (config *Config, err error) {
	config = NewDefaultConfig()
	configPath, _ := os.UserConfigDir()

	// read "config.json"
	data, err := os.ReadFile(filepath.Join(
		configPath, VENDOR_NAME, APPLICATION_NAME, "config.json"))
	if err != nil {
		log.Warn.Printf("read config failed: %s", err.Error())
		return
	}

	// parse json
	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Warn.Printf("parse json config failed: %s", err.Error())
		return
	}

	return
}