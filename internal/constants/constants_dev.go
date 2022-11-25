//go:build dev
// +build dev

package constants

import (
	"encoding/json"
	"os"

	"github.com/knackwurstking/pirgb-web/internal/log"
)

const (
	PIRGB_SERVER_PORT = 50826
)

var (
	ApplicationName   = "pirgb-web"
	VendorName        = "knackwurstking"
)

func LoadConfig() (config *Config, err error) {
	config = NewDefaultConfig()

	// read "config.json"
	data, err := os.ReadFile("config.json")
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
