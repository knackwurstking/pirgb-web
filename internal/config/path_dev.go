//go:build dev
// +build dev

package config

const (
	Vendor     = "knackwurstking"
	Project    = "pirgb-web"
	ConfigFile = "config.dev.yaml"
)

func userConfigDir() string {
	return ""
}
