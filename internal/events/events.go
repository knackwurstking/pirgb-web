package events

import (
	"github.com/knackwurstking/pirgb-web/internal/constants"
	"github.com/knackwurstking/pirgb-web/pkg/pirgb"
)

var (
	Global global
)

// Start is the entry point (will take data from `constants.Config.Devices`)
func Start() {
	var changeEvents []*EventHandler[pirgb.Section]

	for _, device := range constants.Config.Devices {
		for _, section := range device.Sections {
			changeEvents = append(changeEvents, changeEventHandler(device, section))
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

func onChangeEventHandler(event pirgb.Section, device *pirgb.Device, section *pirgb.Section) {
	var pulse int
	var color []int

	for _, pin := range event.Pins {
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
	section.Pins = event.Pins

	go Global.Dispatch(EventNameChange, pirgb.ChangeEvent{
		DeviceEvent: pirgb.DeviceEvent{
			Host: device.Host,
			Port: device.Port,
		},
		Section: pirgb.Section{
			ID:        section.ID,
			Pulse:     section.Pulse,
			LastPulse: section.LastPulse,
			Color:     section.Color,
			Pins:      section.Pins,
		},
	})
}

func changeEventHandler(
	device *pirgb.Device,
	section *pirgb.Section,
) *EventHandler[pirgb.Section] {
	changeEvent := NewChangeEventHandler(device.Host, device.Port, section.ID)
	changeEvent.OnEvent = append(changeEvent.OnEvent, func(event pirgb.Section) {
		onChangeEventHandler(event, device, section)
	})

	return changeEvent
}
