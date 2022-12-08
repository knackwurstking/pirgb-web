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
	ID        int          `json:"id"`
	Pulse     int          `json:"pulse,omitempty"`
	LastPulse int          `json:"lastPulse,omitempty"`
	Color     []int        `json:"color,omitempty"`
	Pins      []SectionPin `json:"pins,omitempty"`
}

type SectionPin struct {
	Pin        int  `json:"pin"`
	Range      int  `json:"range"`
	Pulse      int  `json:"pulse"`
	ColorValue int  `json:"colorValue"`
	ColorPulse int  `json:"colorPulse"`
	IsRunning  bool `json:"isRunning"`
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

// GetPWM data from section for device
func (device *Device) GetPWM(sectionID int) (section *Section, err error) {
	// get pwm section data from device
	resp, err := http.Get(device.URL("pwm", fmt.Sprintf("%d", sectionID)))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return section, &RespError{Resp: resp}
	}

	for _, section = range device.Sections {
		if section.ID == sectionID {
			// parse response
			var sectionData Section
			if err = json.NewDecoder(resp.Body).Decode(&sectionData); err != nil {
				return
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

			break
		}
	}

	if section.ID != sectionID {
		return nil, fmt.Errorf("section \"%d\" for \"%s\" not found", sectionID, device.Host)
	}

	return
}

type Events interface {
	DeviceEvent | ChangeEvent
}

type BaseEventData[T Events] struct {
	Name string `json:"name"`
	Data T      `json:"data"`
}

type DeviceEvent struct { // Used for "offline" and "online" events
	Host string `json:"host"`
	Port int    `json:"port"`
}

type ChangeEvent struct {
	DeviceEvent
	Section
}
