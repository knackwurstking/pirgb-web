package config

import "github.com/sirupsen/logrus"

// Groups handles groups
type Groups []*Group

// Parse sections for all groups
func (groups *Groups) Parse(devices *Devices) {
	for _, group := range *groups {
		group.Parse(devices)
	}
}

// Group ...
type Group struct {
	Name          string   `yaml:"Name"`
	SectionsRegEx []string `yaml:"Sections"`
	Sections      []*Device
}

// Parse sections regex
func (g *Group) Parse(devices *Devices) {
	// TODO: parse sections ...
	logrus.Debugf("... Parsing group \"%s\" [work in progress]", g.Name)
}
