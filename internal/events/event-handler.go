package events

import (
	"context"
	"fmt"
	"sync"
	"time"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"

	"github.com/knackwurstking/pirgb-web/pkg/log"
	"github.com/knackwurstking/pirgb-web/pkg/pirgb"
)

const (
	EventNameChange  = "change"
	EventNameOnline  = "online"
	EventNameOffline = "offline"
)

// EventDataTypes possible for EventHandler
type EventDataTypes interface {
	pirgb.Section
}

// EventHandler handles pirgb-server events like "change"
// (dispatches "change", "online" and "offline" events to Global)
type EventHandler[T EventDataTypes] struct {
	Name      string // "change", ...
	Host      string
	Port      int
	SectionID int
	IsRunning bool
	WaitGroup sync.WaitGroup
	Done      chan struct{}
	Conn      *websocket.Conn
	OnEvent   []func(data T)
}

// NewChangeEventHandler handles "change" server events (when pwm data changes)
func NewChangeEventHandler(
	host string,
	port int,
	sectionID int,
) *EventHandler[pirgb.Section] {
	ev := &EventHandler[pirgb.Section]{
		Name:      EventNameChange,
		Host:      host,
		Port:      port,
		SectionID: sectionID,
		Done:      make(chan struct{}),
		OnEvent:   make([]func(data pirgb.Section), 0),
	}

	return ev
}

func (ev *EventHandler[T]) Connect() error {
	log.Debug.Printf("connect to %s:%d", ev.Host, ev.Port)

	conn, _, err := websocket.Dial(
		context.Background(),
		fmt.Sprintf("ws://%s:%d/ws/event/%s/%d",
			ev.Host, ev.Port, ev.Name, ev.SectionID),
		nil,
	)
	if err != nil {
		return err
	}
	ev.Conn = conn

	Global.Dispatch(EventNameOnline,
		pirgb.DeviceEvent{Host: ev.Host, Port: ev.Port})
	return nil
}

func (ev *EventHandler[T]) reconnect() {
	log.Debug.Printf("reconnect invoked %s:%d", ev.Host, ev.Port)

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

func (ev *EventHandler[T]) Start() {
	if ev.IsRunning {
		return
	}

	log.Debug.Printf("starting event handler %s:%d", ev.Host, ev.Port)
	_ = ev.Connect()

	ev.WaitGroup.Add(1)
	go ev.Handler()
}

func (ev *EventHandler[T]) Stop() {
	log.Debug.Printf("stopping event handler %s:%d", ev.Host, ev.Port)
	ev.Done <- struct{}{}
	ev.WaitGroup.Wait()
}

func (ev *EventHandler[T]) Handler() {
	ev.IsRunning = true

	defer func() {
		ev.IsRunning = false
		ev.Conn.Close(websocket.StatusNormalClosure,
			websocket.StatusNormalClosure.String())
		ev.WaitGroup.Done()
	}()

	go ev.readHandler()

	<-ev.Done
	log.Debug.Printf("EXIT Handler %s:%d", ev.Host, ev.Port)
}

func (ev *EventHandler[T]) readHandler() {
	defer func() {
		ev.Conn.Close(websocket.StatusNormalClosure,
			websocket.StatusNormalClosure.String())

		if !ev.IsRunning { // connection read error
			return
		}

		// dispatch connection closed event to the frontend
		Global.Dispatch(EventNameOffline,
			pirgb.DeviceEvent{Host: ev.Host, Port: ev.Port})

		go ev.reconnect()

		ev.Done <- struct{}{}
	}()

	var err error
	var data T
	for {
		log.Debug.Printf("wait for (JSON) data %s:%d", ev.Host, ev.Port)

		err = wsjson.Read(context.Background(), ev.Conn, &data)
		if err != nil {
			log.Warn.Printf("trouble while reading from device: %s", err.Error())
			break
		}

		// do something
		ev.Dispatch(data)
	}
}

func (ev *EventHandler[T]) Dispatch(data T) {
	log.Debug.Printf("dispatch event %s:%d", ev.Host, ev.Port)

	for _, handler := range ev.OnEvent {
		go handler(data)
	}
}
