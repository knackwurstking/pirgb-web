// Package config handles the server configuration (port, sections, groups, ...)
package config

import (
	"io/ioutil"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

var (
	// GlobalData contains parse configuration data
	GlobalData *Data = &Data{
		Scan:     false,
		Groups:   make(GroupsData, 0),
		Sections: make(SectionsData, 0),
	}
)

func init() {
	// loading configuration first
	file := filepath.Join(userConfigDir(), ConfigFile)
	data, err := ioutil.ReadFile(file)
	if err != nil {
		logrus.Warnf("Read config failed: %s", err.Error())
		return
	}

	err = yaml.Unmarshal(data, GlobalData)
	if err != nil {
		logrus.Warnf("Load config failed: %s", err.Error())
	}
}

// Data is the main config type
type Data struct {
	Scan        bool         `yaml:"Scan"`
	Host        string       `yaml:"Host"`
	Port        int          `yaml:"Port"`
	EnableHTTP  bool         `yaml:"EnableHTTP"`
	EnableHTTPS bool         `yaml:"EnableHTTPS"`
	Groups      GroupsData   `yaml:"Groups"`
	Sections    SectionsData `yaml:"Sections"`
}

// GroupsData handles groups
type GroupsData []GroupData

// GroupData group configuration (name,sections)
type GroupData struct {
	Name     string `yaml:"Name"`
	Sections []int  `yaml:"Sections"`
}

// SectionsData handles sections
type SectionsData []SectionData

// SectionData Section configuration (host,port,sectionID)
type SectionData struct {
	Host      string `yaml:"Host"`
	Port      int    `yaml:"Port"`
	SectionID int    `yaml:"SectionID"`
}
