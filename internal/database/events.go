package database

import (
	"context"
	"fmt"
	"sync"

	"github.com/sirupsen/logrus"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

type Section struct {
	ID   int   `json:"id"`
	Pins []Pin `json:"pins"`
}

type Pin struct {
	Pin        int  `json:"pin"`
	Range      int  `json:"range"`
	Pulse      int  `json:"pulse"`
	ColorValue int  `json:"colorValue"`
	ColorPulse int  `json:"colorPulse"`
	IsRunning  bool `json:"isRunning"`
}

type EventTypes interface {
	Section
}

type Event[T EventTypes] struct {
	Name      string // "change", ...
	IsRunning bool
	WaitGroup *sync.WaitGroup
	Done      chan struct{}
	Conn      *websocket.Conn
	OnEvent   []func(data T)
}

func (ev *Event[T]) Connect() error {
	// TODO: handle reconnects on connection error
	conn, _, err := websocket.Dial(context.Background(), fmt.Sprintf(""), nil)
	if err != nil {
		return err
	}
	ev.Conn = conn
	return nil
}

func (ev *Event[T]) Start() error {
	if ev.IsRunning {
		return nil
	}

	ev.Connect()

	ev.WaitGroup.Add(1)
	go ev.Handler()
	return nil
}

func (ev *Event[T]) Stop() {
	ev.Done <- struct{}{}
	ev.WaitGroup.Wait()
}

func (ev *Event[T]) Dispatch(data T) {
	for _, handler := range ev.OnEvent {
		go handler(data)
	}
}

func (ev *Event[T]) Handler() {
	ev.IsRunning = true
	defer func() {
		ev.IsRunning = false
		ev.WaitGroup.Done()
	}()

	go func() {
		defer func() {
			ev.Conn.Close(websocket.StatusNormalClosure, "bye bye")
			ev.WaitGroup.Done()
			ev.Done <- struct{}{}
		}()

		var err error
		var data T
		for {
			err = wsjson.Read(context.Background(), ev.Conn, &data)
			if err != nil {
				// TODO: check error type, on disconnect to a reconnect every 5 seconds
				logrus.Debugf("change event handler err: %s [type: %T]", err.Error(), err)
				break
			}

			// do something
			ev.Dispatch(data)
		}
	}()

	<-ev.Done
}

func NewChangeEvent() *Event[Section] {
	return &Event[Section]{
		Name:    "change",
		Done:    make(chan struct{}),
		OnEvent: make([]func(data Section), 0),
	}
}
