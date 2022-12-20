package events

import (
	"context"

	"github.com/knackwurstking/pirgb-web/internal/constants"
	"github.com/knackwurstking/pirgb-web/pkg/log"
	"github.com/knackwurstking/pirgb-web/pkg/pirgb"
	"nhooyr.io/websocket"
)

type Client struct {
	Context context.Context
	Conn    *websocket.Conn
	Addr    string
}

type global struct {
	ChangeEvents []*EventHandler[pirgb.Section]
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
	log.Debug.Printf("dispatch \"%s\" event (%#v)", eventName, data)

	switch eventName {
	case EventNameChange:
		dispatchEvent(eventName, pirgb.BaseEvent[pirgb.ChangeEvent]{
			Name: eventName,
			Data: data.(pirgb.ChangeEvent),
		})
	case EventNameOnline, EventNameOffline:
		data := data.(pirgb.DeviceEvent)
		constants.Config.Devices.Get(data.Host).Online = eventName == EventNameOnline
		dispatchEvent(eventName, pirgb.BaseEvent[pirgb.DeviceEvent]{
			Name: eventName,
			Data: data,
		})
	default:
		log.Error.Fatalf("Unknown event name \"%s\"", eventName)
	}
}
