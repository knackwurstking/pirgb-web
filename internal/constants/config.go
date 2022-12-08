package constants

import "github.com/knackwurstking/pirgb-web/pkg/pirgb"

type config struct {
	Port    int           `json:"port"`
	Host    string        `json:"host"`
	Devices pirgb.Devices `json:"devices"`
}

func NewDefaultConfig() *config {
	return &config{
		Port:    50831,
		Host:    "",
		Devices: make(pirgb.Devices, 0),
	}
}
