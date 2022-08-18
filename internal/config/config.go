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
	Host        string  `yaml:"host"`
	Port        int     `yaml:"port"`
	EnableHTTP  bool    `yaml:"enableHTTP"`
	EnableHTTPS bool    `yaml:"enableHTTPS"`
	Devices     Devices `yaml:"devices"`
	Groups      Groups  // no config
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

// NewConfig ...
func NewConfig() *Config {
	c := Config{
		EnableHTTP: true,
		Devices:    make(Devices, 0),
	}
	c.Groups = NewGroups(&c.Devices)

	return &c
}

// DoIt main function to get things running
func DoIt() {
	Global.Load(userConfigDir())

	go func() {
		// scan devices for sections
		logrus.Debugln("Check devices and scan for sections if needed ...")
		Global.Devices.Scan()  // scan if sections are missing in devices
		Global.Devices.Clean() // remove all devices without sections

		// just print out some information
		for _, device := range Global.Devices {
			var sections []int
			for _, section := range device.Sections {
				sections = append(sections, section.ID)
			}

			logrus.WithFields(logrus.Fields{
				"Port":     device.Port,
				"Sections": sections,
				"Groups":   device.Groups,
			}).Debugf("available device: %s", device.Host)
		}
	}()
}
