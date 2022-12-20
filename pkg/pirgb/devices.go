package pirgb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"strings"
	"sync"

	"github.com/knackwurstking/pirgb-web/pkg/log"
)

const (
	PirgbServerPort = 50826
)

// Devices handles device and section related data
type Devices []*Device

func (d Devices) Scan() {
	var wg sync.WaitGroup

	for _, device := range d {
		wg.Add(1)
		go d.ScanDevice(device, &wg)
	}

	wg.Wait()
}

func (d *Devices) ScanDevice(device *Device, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}

	// check and set (default) port
	if device.Port == 0 {
		device.Port = PirgbServerPort
	}

	// get response to "/info" => returns sections `int[]`
	url := fmt.Sprintf("http://%s:%d/info", device.Host, device.Port)
	log.Debug.Printf("scan device \"%s\"", url)

	resp, err := http.Get(url)
	device.Online = err == nil
	if err != nil {
		log.Error.Printf("%s: %s", url, err.Error())
		device.Sections = make([]*Section, 0)
		return
	}
	defer resp.Body.Close()

	// parse response data
	var ids []int
	err = json.NewDecoder(resp.Body).Decode(&ids)
	if err != nil {
		log.Error.Printf("%s: %s", url, err.Error())
	}

	var deviceWG sync.WaitGroup
	devicePWM := func(id int, wg *sync.WaitGroup) {
		defer wg.Done()
		err = device.GetPWM(id)
		if err != nil {
			log.Error.Printf("%s: %s", url, err.Error())
		}
	}

	for _, id := range ids {
		device.Sections = append(device.Sections, &Section{ID: id})
		deviceWG.Add(1)
		go devicePWM(id, &deviceWG)
	}

	deviceWG.Wait()
}

func (d *Devices) Get(host string) *Device {
	for _, device := range *d {
		if device.Host == host {
			return device
		}
	}

	return nil
}

type Device struct {
	Host     string     `json:"host"`
	Port     int        `json:"port"`
	Online   bool       `json:"online"`
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

// GetPWM data from section for device (data will be set to sectionID)
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

	var section *Section
	for _, section = range device.Sections {
		if section.ID == sectionID {
			break
		}

	}

	if section.ID != sectionID {
		return fmt.Errorf("section \"%d\" for \"%s\" not found", sectionID, device.Host)
	}

	// parse response
	var sectionData Section
	if err = json.NewDecoder(resp.Body).Decode(&sectionData); err != nil {
		return err
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
	section.Pins = sectionData.Pins

	return nil
}
