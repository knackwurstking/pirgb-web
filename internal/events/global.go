package events

import (
	"context"
	"time"

	"github.com/knackwurstking/pirgb-web/internal/constants"
	"github.com/knackwurstking/pirgb-web/pkg/log"
	"github.com/knackwurstking/pirgb-web/pkg/pirgb"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
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
	log.Debug.Printf("dispatch \"%s\" event %T", eventName, data)

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

func dispatchEvent[T pirgb.Events](name string, data pirgb.BaseEvent[T]) {
	dispatch := func(client *Client) {
		ctx, cancel := context.WithTimeout(client.Context,
			time.Duration(time.Second*5))
		defer cancel()

		err := wsjson.Write(ctx, client.Conn, data)
		if err == nil {
			return
		}

		// wsjson write error handling
		defer client.Conn.Close(websocket.StatusAbnormalClosure,
			websocket.StatusAbnormalClosure.String())
		log.Warn.Printf("%s: %s [%#v]", name, err, client.Conn)
		Global.RemoveClientAddr(client.Addr)
	}

	for _, client := range Global.Register {
		go dispatch(client)
	}
}
