package events

type Events interface {
	DeviceEventData | ChangeEventData
}

type BaseEventData[T Events] struct {
	Name string `json:"name"`
	Data T      `json:"data"`
}

type DeviceEventData struct { // Used for "offline" and "online" events
	Host string `json:"host"`
	Port int    `json:"port"`
	ID   int    `json:"id" yaml:"id"`
}

type ChangeEventData struct {
	DeviceEventData
	Pulse     int   `json:"pulse" yaml:"pulse"`
	LastPulse int   `json:"lastPulse"`
	Color     []int `json:"color" yaml:"color"`
}
