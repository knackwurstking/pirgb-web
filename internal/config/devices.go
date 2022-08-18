package config

import (
	"sync"

	"github.com/sirupsen/logrus"
	"gitlab.com/knackwurstking/pirgb"
)

// Devices ...
type Devices []*pirgb.Device

// Scan devices for sections available, if sections already set than skip scan
func (*Devices) Scan() {
	var wg sync.WaitGroup

	// scan and parse sections and groups data (using pirgb v0.2.0)
	for index, device := range Global.Devices {
		if len(device.Sections) == 0 {
			wg.Add(1)

			go func(index int, device *pirgb.Device, wg *sync.WaitGroup) {
				defer wg.Done()
				// scan for sections ("/info" endpoint)
				if err := device.Scan(); err != nil || len(device.Sections) == 0 {
					if err != nil {
						logrus.WithField("device", device).Warnln(err.Error())
					} else {
						logrus.WithField("device", device).Warnln("No Sections")
					}
				}
			}(index, device, &wg)
		}
	}

	wg.Wait()
}

// Clean devices, this will remove all devices without sections
func (devices *Devices) Clean() {
	var newDevices []*pirgb.Device
	for _, device := range *devices {
		if len(device.Sections) == 0 {
			continue
		}

		newDevices = append(newDevices, device)
	}
	*devices = newDevices
}

func (devices *Devices) Get(host string) *pirgb.Device {
	for _, device := range *devices {
		if device.Host == host {
			return device
		}
	}

	return nil
}
