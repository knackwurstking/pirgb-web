package constants

import (
	"encoding/json"
	"os"
	"path/filepath"

	"gitlab.com/knackwurstking/pirgb-web/internal/log"
)

const (
	PORT             = 50831
	HOST             = ""
	SERVER_PORT      = 50826
	APPLICATION_NAME = "pirgb-web"
)

func LoadConfig(path string) (config *Config) {
	config = NewDefaultConfig()
	configPath, _ := os.UserConfigDir()

	// read "config.json"
	data, err := os.ReadFile(filepath.Join(configPath, APPLICATION_NAME, "config.json"))
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
