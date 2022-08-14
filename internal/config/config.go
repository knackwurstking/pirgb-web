// Package config handles the server configuration (port, sections, groups, ...)
package config

import (
	"io/ioutil"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

// Config is the main config type
type Config struct {
	Scan        bool    `yaml:"Scan"`
	Host        string  `yaml:"Host"`
	Port        int     `yaml:"Port"`
	EnableHTTP  bool    `yaml:"EnableHTTP"`
	EnableHTTPS bool    `yaml:"EnableHTTPS"`
	Groups      Groups  `yaml:"Groups"`
	Devices     Devices `yaml:"Devices"`
}

// Load configuration
func (config *Config) Load(path string) {
	// loading configuration first
	file := filepath.Join(path, ConfigFile)
	data, err := ioutil.ReadFile(file)
	if err != nil {
		logrus.Warnf("Read config failed: %s", err.Error())
	} else {
		// merge config into `Global` (if possible)
		err = yaml.Unmarshal(data, config)
		if err != nil {
			logrus.Warnf("Load config failed: %s", err.Error())
		}
	}
}

// DoIt main function to get things running
func DoIt() {
	Global.Load(userConfigDir())

	go func() {
		// scan devices for sections
		logrus.Debugln("check devices and scan for sections if needed ...")
		Global.Devices.Scan()  // scan if sections are missing in devices
		Global.Devices.Clean() // remove all devices without sections

		// parse groups
		logrus.Debugln("sections scan done, parse groups now ...")
		Global.Groups.Parse(&Global.Devices)
	}()
}
