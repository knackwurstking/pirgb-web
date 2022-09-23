package events

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"gitlab.com/knackwurstking/pirgb-web/internal/servertypes"
	"gitlab.com/knackwurstking/pirgb-web/pkg/pirgb"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

type EventHandler[T servertypes.EventTypes] struct {
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

func (ev *EventHandler[T]) Connect() error {
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

	Global.Dispatch("online", pirgb.DeviceEventData{
		Host: ev.Host,
		Port: ev.Port,
	})

	return nil
}

func (ev *EventHandler[T]) reconnect() {
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

func (ev *EventHandler[T]) Start() error {
	if ev.IsRunning {
		return nil
	}

	ev.Log.Debugf("[events] starting event handler")

	ev.Connect()

	ev.WaitGroup.Add(1)
	go ev.Handler()
	return nil
}

func (ev *EventHandler[T]) Stop() {
	ev.Log.Debugf("[events] stopping event handler")
	ev.Done <- struct{}{}
	ev.WaitGroup.Wait()
}

func (ev *EventHandler[T]) Handler() {
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
				Global.Dispatch("offline", pirgb.DeviceEventData{
					Host: ev.Host,
					Port: ev.Port,
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

func (ev *EventHandler[T]) Dispatch(data T) {
	ev.Log.Debugf("[events] dispatch event")

	for _, handler := range ev.OnEvent {
		go handler(data)
	}
}

func NewChangeEventHandler(host string, port int, sectionID int) *EventHandler[servertypes.Section] {
	ev := &EventHandler[servertypes.Section]{
		Name:      "change",
		Host:      host,
		Port:      port,
		SectionID: sectionID,
		Done:      make(chan struct{}),
		OnEvent:   make([]func(data servertypes.Section), 0),
	}

	ev.Log = logrus.WithFields(logrus.Fields{
		"Host":      ev.Host,
		"Port":      ev.Port,
		"SectionID": ev.SectionID,
	})

	return ev
}
