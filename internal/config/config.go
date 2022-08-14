// Package config handles the server configuration (port, sections, groups, ...)
package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"sync"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

const (
	DefaultPort = 50826
)

var (
	// GlobalData contains parse configuration data
	GlobalData = &Data{
		Scan:     false,
		Groups:   make(GroupsData, 0),
		Sections: make(SectionsData, 0),
	}
)

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
	Name          string   `yaml:"Name"`
	SectionsRegEx []string `yaml:"Sections"`
	Sections      []*SectionData
}

// Parse ...
func (group *GroupData) Parse(sectionsData *SectionsData) {
	// TODO: parse section (regex) and store section from SectionsData in Sections
}

// SectionsData handles sections
type SectionsData []SectionData

// SectionData Section configuration (host,port,sectionID)
type SectionData struct {
	Host     string `yaml:"Host"`
	Port     int    `yaml:"Port"`
	Sections []int  `yaml:"Sections"`
}

// Scan for sections if `Sections` field is empty
func (s *SectionData) Scan() error {
	r, err := http.Get(fmt.Sprintf("http://%s:%d/info", s.Host, s.Port))
	if err != nil {
		return err
	}
	defer r.Body.Close()

	err = json.NewDecoder(r.Body).Decode(&s.Sections)
	if err != nil {
		return err
	}

	return nil
}

// DoIt main function to get things running
func DoIt() {
	// loading configuration first
	file := filepath.Join(userConfigDir(), ConfigFile)
	data, err := ioutil.ReadFile(file)
	if err != nil {
		logrus.Warnf("Read config failed: %s", err.Error())
	} else {
		// merge config into `GlobalData` (if possible)
		err = yaml.Unmarshal(data, GlobalData)
		if err != nil {
			logrus.Warnf("Load config failed: %s", err.Error())
		}
	}

	go func() {
		var wg sync.WaitGroup

		// scan and parse sections and groups data (using pirgb v0.2.0)
		for _, section := range GlobalData.Sections {
			if len(section.Sections) == 0 {
				wg.Add(1)
				go func(wg *sync.WaitGroup) {
					defer wg.Done()
					// scan for sections ("/info" endpoint)
					if err := section.Scan(); err != nil {
						// TODO: remove section from the list and continue (and log)
						// ...
					} else if len(section.Sections) == 0 {
						// TODO: if no sections found also remove section from the list (and log)
						// ...
					}
				}(&wg)
			}
		}

		wg.Wait()
		// TODO: parse groups if scan has finished
		// ...
	}()
}
