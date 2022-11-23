package constants

import "gitlab.com/knackwurstking/pirgb-web/pkg/pirgb"

type Config struct {
	Port    int     `json:"port"`
	Host    string  `json:"host"`
	Devices Devices `json:"devices"`
}

func NewDefaultConfig() *Config {
	return &Config{
		Port:    50831,
		Host:    "",
		Devices: make(Devices, 0),
	}
}

type Devices []*pirgb.Device
