package config

// Groups ...
type Groups struct {
	Devices *Devices
}

// TODO: list groups available: []string
// TODO: get group: returns a new filtered list of devices []*Devices

// NewGroups ...
func NewGroups(devices *Devices) Groups {
	return Groups{
		Devices: devices,
	}
}
