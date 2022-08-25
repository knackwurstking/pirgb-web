package events

// TODO: Clean up this mess here

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"gitlab.com/knackwurstking/pirgb-web/internal/config"
	"gitlab.com/knackwurstking/pirgb-web/pkg/pirgb"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

var (
	Global global
)

type ChangeEventData struct {
	Host string `json:"host"`
	Port int    `json:"port"`

	ID        int   `json:"id" yaml:"id"`
	Pulse     int   `json:"pulse" yaml:"pulse"`
	LastPulse int   `json:"lastPulse"`
	Color     []int `json:"color" yaml:"color"`
}

type DeviceStateEventData struct { // Used for "offline" and "online" events
	Host string `json:"host"`
	Port int    `json:"port"`
	ID   int    `json:"id" yaml:"id"`
}

type Client struct {
	Context context.Context
	Conn    *websocket.Conn
	Addr    string
}

type global struct {
	ChangeEvents []*Event[pirgb.Section]
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

func (g *global) Dispatch(eventName string, data any) {
	logrus.Debugf("[events] dispatch \"%s\" event", eventName)

	switch eventName {
	case "change":
		for _, client := range g.Register {
			go func(client *Client) {
				ctx, cancel := context.WithTimeout(client.Context, time.Duration(time.Second*5))
				defer cancel()

				err := wsjson.Write(ctx, client.Conn, struct {
					Name string          `json:"name"`
					Data ChangeEventData `json:"data"`
				}{Name: eventName, Data: data.(ChangeEventData)})

				if err != nil {
					logrus.WithFields(logrus.Fields{
						"eventName": eventName,
						"conn":      client.Conn,
					}).Warnln("[Events]", err.Error())

					// Remove client address from register
					client.Conn.Close(websocket.StatusAbnormalClosure, "read failed, close connection, remove client from register")
					g.RemoveClientAddr(client.Addr)
				}
			}(client)
		}
	case "offline", "online":
		for _, client := range g.Register {
			go func(client *Client) {
				ctx, cancel := context.WithTimeout(client.Context, time.Duration(time.Second*5))
				defer cancel()

				err := wsjson.Write(ctx, client.Conn, struct {
					Name string               `json:"name"`
					Data DeviceStateEventData `json:"data"`
				}{Name: eventName, Data: data.(DeviceStateEventData)})

				if err != nil {
					logrus.WithFields(logrus.Fields{
						"eventName": eventName,
						"conn":      client.Conn,
					}).Warnln("[Events]", err.Error())

					// Remove client address from register
					client.Conn.Close(websocket.StatusAbnormalClosure, "read failed, close connection, remove client from register")
					g.RemoveClientAddr(client.Addr)
				}
			}(client)
		}
	default:
		logrus.Fatalf("Unknown event name \"%s\"", eventName)
	}
}

type Event[T pirgb.EventTypes] struct {
	Name      string // "change", ...
	Host      string
	Port      int
	SectionID int
	IsRunning bool
	WaitGroup sync.WaitGroup
	Done      chan struct{}
	Conn      *websocket.Conn
	OnEvent   []func(data T)
	Log       *logrus.Entry
}

func (ev *Event[T]) Connect() error {
	ev.Log.Debugf("[events] Try to connect...")
	conn, _, err := websocket.Dial(
		context.Background(),
		fmt.Sprintf("ws://%s:%d/ws/event/%s/%d", ev.Host, ev.Port, ev.Name, ev.SectionID),
		nil,
	)
	if err != nil {
		return err
	}
	ev.Conn = conn

	Global.Dispatch("online", DeviceStateEventData{
		Host: ev.Host,
		Port: ev.Port,
		ID:   ev.SectionID,
	})

	return nil
}

func (ev *Event[T]) reconnect() {
	ev.Log.Debugf("[events] reconnect invoked...")
	var err error

	for {
		time.Sleep(time.Millisecond * 2500)
		err = ev.Connect()
		if err == nil {
			break
		}
	}

	ev.WaitGroup.Wait()
	ev.WaitGroup.Add(1)
	go ev.Handler()
}

func (ev *Event[T]) Start() error {
	if ev.IsRunning {
		return nil
	}

	ev.Log.Debugf("[events] starting event handler")

	ev.Connect()

	ev.WaitGroup.Add(1)
	go ev.Handler()
	return nil
}

func (ev *Event[T]) Stop() {
	ev.Log.Debugf("[events] stopping event handler")
	ev.Done <- struct{}{}
	ev.WaitGroup.Wait()
}

func (ev *Event[T]) Handler() {
	ev.IsRunning = true
	defer func() {
		ev.IsRunning = false
		ev.Conn.Close(websocket.StatusNormalClosure, "bye bye")
		ev.WaitGroup.Done()
	}()

	go func() {
		defer func() {
			ev.Conn.Close(websocket.StatusNormalClosure, "bye bye")

			if ev.IsRunning { // connection read error
				// dispatch connection closed event to the frontend
				Global.Dispatch("offline", DeviceStateEventData{
					Host: ev.Host,
					Port: ev.Port,
					ID:   ev.SectionID,
				})

				go ev.reconnect()
				ev.Done <- struct{}{}
			}
		}()

		var err error
		var data T
		for {
			ev.Log.Debugf("[events] wait for (json) data")

			err = wsjson.Read(context.Background(), ev.Conn, &data)
			if err != nil {
				ev.Log.Warnf("[events] trouble while reading from device: %s", err.Error())
				ev.Log.Debugf("[events] error type: %T", err)
				break
			}

			// do something
			ev.Dispatch(data)
		}
	}()

	<-ev.Done
	ev.Log.Debugf("[events] EXIT Handler")
}

func (ev *Event[T]) Dispatch(data T) {
	ev.Log.Debugf("[events] dispatch event")

	for _, handler := range ev.OnEvent {
		go handler(data)
	}
}

func NewChangeEvent(host string, port int, sectionID int) *Event[pirgb.Section] {
	ev := &Event[pirgb.Section]{
		Name:      "change",
		Host:      host,
		Port:      port,
		SectionID: sectionID,
		Done:      make(chan struct{}),
		OnEvent:   make([]func(data pirgb.Section), 0),
	}

	ev.Log = logrus.WithFields(logrus.Fields{
		"Host":      ev.Host,
		"Port":      ev.Port,
		"SectionID": ev.SectionID,
	})

	return ev
}

func Initialize() {
	var changeEvents []*Event[pirgb.Section]

	for _, device := range config.Global.Devices {
		for _, section := range device.Sections {
			func(device *config.Device, section *config.Section) {
				changeEvent := NewChangeEvent(device.Host, device.Port, section.ID)
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

					go Global.Dispatch(changeEvent.Name, ChangeEventData{
						Host:      device.Host,
						Port:      device.Port,
						ID:        section.ID,
						Pulse:     section.Pulse,
						LastPulse: section.LastPulse,
						Color:     section.Color,
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
