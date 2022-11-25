package events

import (
	"context"
	"fmt"
	"sync"
	"time"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"

	"github.com/knackwurstking/pirgb-web/internal/log"
	"github.com/knackwurstking/pirgb-web/pkg/pirgb"
)

type EventHandler[T EventTypes] struct {
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

func (ev *EventHandler[T]) Connect() error {
	log.Debug.Println("try to connect...")
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

	Global.Dispatch("online", pirgb.DeviceEventData{
		Host: ev.Host,
		Port: ev.Port,
	})

	return nil
}

func (ev *EventHandler[T]) reconnect() {
	log.Debug.Println("reconnect invoked...")
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

	log.Debug.Println("starting event handler")
	ev.Connect()

	ev.WaitGroup.Add(1)
	go ev.Handler()
	return nil
}

func (ev *EventHandler[T]) Stop() {
	log.Debug.Println("stopping event handler")
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
			log.Debug.Println("wait for (json) data")

			err = wsjson.Read(context.Background(), ev.Conn, &data)
			if err != nil {
				log.Warn.Printf("trouble while reading from device: %s", err.Error())
				log.Debug.Printf("error type: %T", err)
				break
			}

			// do something
			ev.Dispatch(data)
		}
	}()

	<-ev.Done
	log.Debug.Println("EXIT Handler")
}

func (ev *EventHandler[T]) Dispatch(data T) {
	log.Debug.Println("dispatch event")

	for _, handler := range ev.OnEvent {
		go handler(data)
	}
}

func NewChangeEventHandler(host string, port int, sectionID int) *EventHandler[Section] {
	ev := &EventHandler[Section]{
		Name:      "change",
		Host:      host,
		Port:      port,
		SectionID: sectionID,
		Done:      make(chan struct{}),
		OnEvent:   make([]func(data Section), 0),
	}

	return ev
}
