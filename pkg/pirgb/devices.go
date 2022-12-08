package pirgb

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/knackwurstking/pirgb-web/pkg/log"
)

const (
	PIRGB_SERVER_PORT = 50826
)

// Devices handles device and section related data
type Devices []*Device

func (d Devices) Scan() {
	var wg sync.WaitGroup

	for _, device := range d {
		wg.Add(1)
		go func(device *Device, wg *sync.WaitGroup) {
			defer wg.Done()

			// check and set (default) port
			if device.Port == 0 {
				device.Port = PIRGB_SERVER_PORT
			}

			// get response to "/info" => returns sections `int[]`
			url := fmt.Sprintf("http://%s:%d/info", device.Host, device.Port)
			log.Debug.Printf("scan device \"%s\"", url)

			resp, err := http.Get(url)
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
			for _, id := range ids {
				device.Sections = append(device.Sections, &Section{ID: id})

				deviceWG.Add(1)
				// init: device data for section <id>
				go func(id int, wg *sync.WaitGroup) {
					defer wg.Done()
					_, err = device.GetPWM(id)
					if err != nil {
						log.Error.Printf("%s: %s", url, err.Error())
					}
				}(id, &deviceWG)
			}
			deviceWG.Wait()
		}(device, &wg)
	}

	wg.Wait()
}

func (d *Devices) Get(host string) *Device {
	for _, device := range *d {
		if device.Host == host {
			return device
		}
	}

	return nil
}
