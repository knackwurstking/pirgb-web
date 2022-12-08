package constants

import "github.com/knackwurstking/pirgb-web/pkg/pirgb"

type Config struct {
	Port    int           `json:"port"`
	Host    string        `json:"host"`
	Devices pirgb.Devices `json:"devices"`
}

func NewDefaultConfig() *Config {
	return &Config{
		Port:    50831,
		Host:    "",
		Devices: make(pirgb.Devices, 0),
	}
}
