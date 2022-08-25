package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"
	"gitlab.com/knackwurstking/pirgb-web/pkg/pirgb"
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
			logrus.WithField("url", url).Debugln("[config] scan device")

			resp, err := http.Get(url)
			if err != nil {
				logrus.WithField("url", url).Errorln("[config]", err.Error())
				device.Sections = make([]*Section, 0)
				return
			}

			defer resp.Body.Close()

			var sectionIDs []int
			err = json.NewDecoder(resp.Body).Decode(&sectionIDs)
			if err != nil {
				logrus.WithField("url", url).Errorln("[config]", err.Error())
			}

			for _, sectionID := range sectionIDs {
				device.Sections = append(device.Sections, &Section{ID: sectionID})
				go device.GetPWM(sectionID)
			}
		}(device, &wg)
	}

	wg.Wait()
}

func (devices *Devices) Get(host string) *Device {
	for _, device := range *devices {
		if device.Host == host {
			return device
		}
	}

	return nil
}

type Section struct {
	ID        int   `json:"id" yaml:"id"`
	Pulse     int   `json:"pulse" yaml:"pulse"`
	LastPulse int   `json:"lastPulse"`
	Color     []int `json:"color" yaml:"color"`
}

type Device struct {
	Host     string     `json:"host" yaml:"host"`
	Port     int        `json:"port" yaml:"port"`
	Sections []*Section `json:"sections" yaml:"sections"`
	Groups   []string   `json:"groups" yaml:"groups"`
}

func (device *Device) GetSection(sectionID int) *Section {
	for _, section := range device.Sections {
		if section.ID == sectionID {
			return section
		}
	}

	return nil
}

func (device *Device) URL(path ...string) string {
	return fmt.Sprintf(
		"http://%s:%d/%s",
		device.Host, device.Port, strings.TrimLeft(filepath.Join(path...), "/"),
	)
}

func (device *Device) SetPWM(sectionID int, pwm pirgb.PWM) error {
	url := device.URL("pwm", fmt.Sprintf("%d", sectionID))
	data, _ := json.Marshal(pwm)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return &RespError{Resp: resp}
	}

	return nil
}

func (device *Device) GetPWM(sectionID int) error {
	// get pwm section data from device
	resp, err := http.Get(device.URL("pwm", fmt.Sprintf("%d", sectionID)))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return &RespError{Resp: resp}
	}

	for _, section := range device.Sections {
		if section.ID == sectionID {
			// parse response
			var sectionData struct {
				ID   int `json:"id"`
				Pins []struct {
					Pulse      int `json:"pulse"`
					ColorValue int `json:"colorValue"`
				} `json:"pins"`
			}
			if err := json.NewDecoder(resp.Body).Decode(&sectionData); err != nil {
				return nil
			}

			// parse pulse and color
			var pulse int
			var color []int
			for _, pin := range sectionData.Pins {
				if pin.Pulse > 0 {
					pulse = pin.Pulse
				}

				color = append(color, pin.ColorValue)
			}

			// set data to section
			section.Pulse = pulse
			section.Color = color
		}
	}

	return nil
}
