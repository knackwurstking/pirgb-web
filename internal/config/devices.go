package config

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/sirupsen/logrus"
)

// Devices ...
type Devices []Device

// Scan devices for sections available, if sections already set than skip scan
func (devices *Devices) Scan() {
	var wg sync.WaitGroup

	// scan and parse sections and groups data (using pirgb v0.2.0)
	for index, device := range Global.Devices {
		if len(device.Sections) == 0 {
			wg.Add(1)

			go func(index int, device *Device, wg *sync.WaitGroup) {
				defer wg.Done()
				// scan for sections ("/info" endpoint)
				if err := device.Scan(); err != nil || len(device.Sections) == 0 {
					if err != nil {
						logrus.WithField("device", device).Warnln(err.Error())
					} else {
						logrus.WithField("device", device).Warnln("No Sections")
					}
				}
			}(index, &device, &wg)
		}
	}

	wg.Wait()
}

// Clean devices, this will remove all devices without sections
func (devices *Devices) Clean() {
	var newDevices []Device
	for _, device := range *devices {
		if len(device.Sections) == 0 {
			continue
		}

		newDevices = append(newDevices, device)
	}
	*devices = newDevices
}

// Device ...
type Device struct {
	Host     string `yaml:"Host"`
	Port     int    `yaml:"Port"`
	Sections []int  `yaml:"Sections"` // just the section ids
}

// Scan for sections if `Sections` field is empty
func (s *Device) Scan() error {
	if s.Port == 0 {
		s.Port = DefaultPort
	}
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
