package config

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/sirupsen/logrus"
	"gitlab.com/knackwurstking/pirgb-web/pkg/pirgb"
)

// Devices ...
type Devices []*pirgb.Device

// Scan devices for sections available, if sections already set than skip scan
func (devices *Devices) Scan() {
	var wg sync.WaitGroup

	// scan and parse sections and groups data (using pirgb v0.2.0)
	for _, device := range *devices {
		wg.Add(1)
		func(device *pirgb.Device, wg *sync.WaitGroup) {
			defer wg.Done()

			if device.Port == 0 {
				device.Port = DefaultServerPort
			}

			url := fmt.Sprintf("http://%s:%d/info", device.Host, device.Port)
			logrus.WithField("url", url).Debugln("[config] scan device")

			resp, err := http.Get(url)
			if err != nil {
				logrus.WithField("url", url).Errorln("[config]", err.Error())
				device.Sections = make([]*pirgb.Section, 0)
				return
			}

			defer resp.Body.Close()

			var sectionIDs []int
			err = json.NewDecoder(resp.Body).Decode(&sectionIDs)
			if err != nil {
				logrus.WithField("url", url).Errorln("[config]", err.Error())
			}

			for _, sectionID := range sectionIDs {
				device.Sections = append(device.Sections, &pirgb.Section{ID: sectionID})
				go device.GetPWM(sectionID)
			}
		}(device, &wg)
	}

	wg.Wait()
}

func (devices *Devices) Get(host string) *pirgb.Device {
	for _, device := range *devices {
		if device.Host == host {
			return device
		}
	}

	return nil
}
