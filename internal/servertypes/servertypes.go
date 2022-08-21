package servertypes

// PWM data to send to pirgb-server (pwm request)
type PWM struct {
	Pulse int   `json:"pulse"`
	RGBW  []int `json:"rgbw"`
}

// Section returned from pirgb-server (pwm request)
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
