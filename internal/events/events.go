package events

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"

	"gitlab.com/knackwurstking/pirgb-web/internal/constants"
	"gitlab.com/knackwurstking/pirgb-web/pkg/pirgb"
)

var (
	Global global
)

type Client struct {
	Context context.Context
	Conn    *websocket.Conn
	Addr    string
}

type global struct {
	ChangeEvents []*EventHandler[Section]
	Register     []*Client
}

// AddClient to the register
func (g *global) AddClient(ctx context.Context, conn *websocket.Conn, addr string) {
	for _, client := range g.Register {
		if client.Addr == addr {
			client.Conn.Close(websocket.StatusAbnormalClosure, "recreate connection")
			client.Conn = conn
			client.Context = ctx
			return
		}
	}

	g.Register = append(g.Register, &Client{
		Context: ctx,
		Conn:    conn,
		Addr:    addr,
	})
}

// RemoveClientAddr from the register
func (g *global) RemoveClientAddr(addr string) {
	var newRegister []*Client
	for _, client := range g.Register {
		if client.Addr != addr {
			newRegister = append(newRegister, client)
		}
	}
	g.Register = newRegister
}

// Dispatch event with data ("change"|"online"|"offline")
func (g *global) Dispatch(eventName string, data any) {
	logrus.Debugf("[events] dispatch \"%s\" event", eventName)

	switch eventName {
	case "change":
		dispathEvent(
			g, eventName,
			pirgb.BaseEventData[pirgb.ChangeEventData]{
				Name: eventName,
				Data: data.(pirgb.ChangeEventData),
			},
		)
	case "offline", "online":
		dispathEvent(
			g, eventName,
			pirgb.BaseEventData[pirgb.DeviceEventData]{
				Name: eventName,
				Data: data.(pirgb.DeviceEventData),
			},
		)
	default:
		logrus.Fatalf("Unknown event name \"%s\"", eventName)
	}
}

func dispathEvent[T pirgb.Events](g *global, name string, data pirgb.BaseEventData[T]) {
	for _, client := range g.Register {
		go func(client *Client) {
			ctx, cancel := context.WithTimeout(client.Context, time.Duration(time.Second*5))
			defer cancel()

			err := wsjson.Write(ctx, client.Conn, data)

			if err != nil {
				logrus.WithFields(logrus.Fields{
					"eventName": name,
					"conn":      client.Conn,
				}).Warnln("[Events]", err.Error())

				// Remove client address from register
				client.Conn.Close(websocket.StatusAbnormalClosure, "read failed, close connection, remove client from register")
				g.RemoveClientAddr(client.Addr)
			}
		}(client)
	}
}

func Initialize(config *constants.Config) {
	var changeEvents []*EventHandler[Section]

	for _, device := range config.Devices {
		for _, section := range device.Sections {
			func(device *pirgb.Device, section *pirgb.Section) {
				changeEvent := NewChangeEventHandler(device.Host, device.Port, section.ID)
				changeEvent.OnEvent = append(changeEvent.OnEvent, func(data Section) {
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

					go Global.Dispatch(changeEvent.Name, pirgb.ChangeEventData{
						DeviceEventData: pirgb.DeviceEventData{
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
