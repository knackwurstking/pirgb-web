package pirgb

import (
	"fmt"
	"net/http"
)

type RespError struct {
	Resp *http.Response
}

func (err *RespError) Error() string {
	return fmt.Sprintf("[%s] %s", err.Resp.Status, err.Resp.Request.URL)
}

type PWM struct {
	Pulse int   `json:"pulse"`
	Color []int `json:"color"`
}

type Section struct {
	ID        int    `json:"id"`
	Pulse     int    `json:"pulse,omitempty"`
	LastPulse int    `json:"lastPulse,omitempty"`
	Color     []int  `json:"color,omitempty"`
	Pins      []*Pin `json:"pins,omitempty"`
}

type Pin struct {
	Pin        int  `json:"pin"`
	Range      int  `json:"range"`
	Pulse      int  `json:"pulse"`
	ColorValue int  `json:"colorValue"`
	ColorPulse int  `json:"colorPulse"`
	IsRunning  bool `json:"isRunning"`
}

type Events interface {
	DeviceEvent | ChangeEvent
}

// BaseEvent struct used for dispatch event to frontend
type BaseEvent[T Events] struct {
	Name string `json:"name"`
	Data T      `json:"data"`
}

// DeviceEvent - device is "online", "offline"
type DeviceEvent struct { // Used for "offline" and "online" events
	Host   string `json:"host"`
	Port   int    `json:"port"`
	Online bool   `json:"online"`
}

// ChangeEvent - device (section) pwm state changed
type ChangeEvent struct {
	DeviceEvent
	Section
}
