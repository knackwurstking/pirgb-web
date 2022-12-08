package events

import (
	"context"
	"time"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"

	"github.com/knackwurstking/pirgb-web/internal/constants"
	"github.com/knackwurstking/pirgb-web/pkg/log"
	"github.com/knackwurstking/pirgb-web/pkg/pirgb"
)

var (
	Global global
)

func dispathEvent[T pirgb.Events](name string, data pirgb.BaseEvent[T]) {
	for _, client := range Global.Register {
		go func(client *Client) {
			ctx, cancel := context.WithTimeout(client.Context, time.Duration(time.Second*5))
			defer cancel()

			err := wsjson.Write(ctx, client.Conn, data)

			if err != nil {
				log.Warn.Printf("%s: %s [%+v]", name, err, client.Conn)

				// Remove client address from register
				client.Conn.Close(websocket.StatusAbnormalClosure, "read failed, close connection, remove client from register")
				Global.RemoveClientAddr(client.Addr)
			}
		}(client)
	}
}

func Start() {
	var changeEvents []*EventHandler[pirgb.Section]

	for _, device := range constants.Config.Devices {
		for _, section := range device.Sections {
			func(device *pirgb.Device, section *pirgb.Section) {
				changeEvent := NewChangeEventHandler(device.Host, device.Port, section.ID)
				changeEvent.OnEvent = append(changeEvent.OnEvent, func(data pirgb.Section) {
					var pulse int
					var color []int

					for _, pin := range data.Pins {
						if pin.Pulse > 0 {
							pulse = pin.Pulse
						}

						color = append(color, pin.ColorValue)
					}

					if section.Pulse > 0 {
						section.LastPulse = section.Pulse
					}
					section.Pulse = pulse
					section.Color = color

					go Global.Dispatch(changeEvent.Name, pirgb.ChangeEvent{
						DeviceEvent: pirgb.DeviceEvent{
							Host: device.Host,
							Port: device.Port,
						},
						Section: pirgb.Section{
							ID:        section.ID,
							Pulse:     section.Pulse,
							LastPulse: section.LastPulse,
							Color:     section.Color,
						},
					})
				})
				changeEvents = append(changeEvents, changeEvent)
			}(device, section)
		}
	}

	// (re)start event handler for devices
	for _, changeEvent := range Global.ChangeEvents {
		changeEvent.Stop()
	}

	Global.ChangeEvents = changeEvents

	for _, changeEvent := range Global.ChangeEvents {
		changeEvent.Start()
	}
}
