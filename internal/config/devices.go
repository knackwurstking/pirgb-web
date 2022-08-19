package config

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/sirupsen/logrus"
)

// Devices ...
type Devices []*Device

// Scan devices for sections available, if sections already set than skip scan
func (devices *Devices) Scan() {
	var wg sync.WaitGroup

	// scan and parse sections and groups data (using pirgb v0.2.0)
	for _, device := range *devices {
		wg.Add(1)
		func(device *Device, wg *sync.WaitGroup) {
			defer wg.Done()

			if device.Port == 0 {
				device.Port = DefaultServerPort
			}

			url := fmt.Sprintf("http://%s:%d/info", device.Host, device.Port)
			logrus.WithField("url", url).Debugln("scan device")

			resp, err := http.Get(url)
			if err != nil {
				logrus.WithField("url", url).Errorln(err.Error())
				device.Sections = make([]int, 0)
				return
			}

			defer resp.Body.Close()

			err = json.NewDecoder(resp.Body).Decode(&device.Sections)
			if err != nil {
				logrus.WithField("url", url).Errorln(err.Error())
			}
		}(device, &wg)
	}

	wg.Wait()
}

type Device struct {
	Host     string   `json:"host" yaml:"host"`
	Port     int      `json:"port" yaml:"port"`
	Sections []int    `json:"sections" yaml:"sections"`
	Groups   []string `json:"groups" yaml:"groups"`
}
