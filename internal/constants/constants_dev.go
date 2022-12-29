//go:build dev
// +build dev

package constants

import (
	"encoding/json"
	"os"

	"github.com/knackwurstking/pirgb-web/internal/middleware"
	"github.com/knackwurstking/pirgb-web/pkg/log"
)

const (
	PIRGB_SERVER_PORT = 50826
)

var (
	ApplicationName string = "pirgb-web"
	VendorName      string = "knackwurstking"
	Config          *config
	User            *middleware.User = &middleware.User{
		Name:     os.Getenv("AUTH_USERNAME"),
		Password: os.Getenv("AUTH_PASSWORD"),
	}
)

func LoadConfig() error {
	Config = NewDefaultConfig()

	// read "config.json"
	data, err := os.ReadFile("config.json")
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
