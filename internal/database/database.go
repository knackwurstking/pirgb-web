package database

// What do i need here?
//	- list devices (host, port, sections[int])
//	- store device and section specific data
//		- id (primary key for sections)
//		- pulse (int) (always bigger 0)
//		- lastPulse (int) (always bigger 0)
//		- rgbw ([4]int)
//	- Event handler (pirgb-server "change" event)

type Pulse int

type Color []int

type Devices struct {
	List []DeviceSection
}

type DeviceSection struct {
	ID        int
	Pulse     int
	LastPulse int
	Color     Color
}

func Initialize() {
	// TODO: start event handler
}
