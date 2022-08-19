// Package config handles the server configuration (port, sections, groups, ...)
package config

import (
	"io/ioutil"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

const (
	DefaultServerWebPort = 50831
	DefaultServerPort    = 50826
)

var (
	Global *GlobalConfig
)

func init() {
	Global = NewGlobalConfig(GetConfigPath())
}

type GlobalConfig struct {
	Path    string  `json:"-"`                      // empty string if configuration not loaded
	Host    string  `json:"host" yaml:"host"`       // server host address (`pirgb-web`)
	Port    int     `json:"port" yaml:"port"`       // server port (`pirgb-web`)
	Devices Devices `json:"devices" yaml:"devices"` // list of devices: (host, port, section ids)
}

func (config *GlobalConfig) Load() {
	// loading configuration first
	data, err := ioutil.ReadFile(filepath.Join(config.Path, configFile))
	if err != nil {
		logrus.Warnf("Read config failed: %s", err.Error())
		return
	}

	// merge config into `Global` (if possible)
	err = yaml.Unmarshal(data, config)
	if err != nil {
		logrus.Warnf("Load config failed: %s", err.Error())
	}

	config.Devices.Scan()
}

func NewGlobalConfig(path string) *GlobalConfig {
	c := GlobalConfig{
		Path:    path,
		Port:    DefaultServerWebPort,
		Devices: make(Devices, 0),
	}
	c.Load()

	return &c
}
