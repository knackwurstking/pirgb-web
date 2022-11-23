package pirgb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"strings"
)

type RespError struct {
	Resp *http.Response
}

func (err *RespError) Error() string {
	return fmt.Sprintf("[%s] %s", err.Resp.Status, err.Resp.Request.URL)
}

type PWM struct {
	Pulse int   `json:"pulse"`
	Color []int `json:"color"`
}

type Section struct {
	ID        int   `json:"id"`
	Pulse     int   `json:"pulse"`
	LastPulse int   `json:"lastPulse"`
	Color     []int `json:"color"`
}

type Device struct {
	Host     string     `json:"host"`
	Port     int        `json:"port"`
	Sections []*Section `json:"sections"`
	Groups   []string   `json:"groups"`
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

func (device *Device) SetPWM(sectionID int, pwm PWM) error {
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

type Events interface {
	DeviceEventData | ChangeEventData
}

type BaseEventData[T Events] struct {
	Name string `json:"name"`
	Data T      `json:"data"`
}

type DeviceEventData struct { // Used for "offline" and "online" events
	Host string `json:"host"`
	Port int    `json:"port"`
}

type ChangeEventData struct {
	DeviceEventData
	Section
}
