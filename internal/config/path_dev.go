//go:build dev
// +build dev

package config

const (
	Vendor     = "knackwurstking"
	Project    = "pirgb-web"
	configFile = "config.dev.yaml"
)

func GetConfigPath() string {
	return ""
}
