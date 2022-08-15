package config

// Groups ...
type Groups struct {
	Devices *Devices `json:"-"`
}

// ListGroups for all devices
func (g *Groups) ListGroups() []string {
	var groups []string

	for _, device := range *g.Devices {
		for _, group := range device.Groups {
			// check if group already in groups
			var exists bool

			for _, g := range groups {
				if g == group {
					exists = true
					break
				}
			}

			if !exists {
				groups = append(groups, group)
			}
		}
	}

	return groups
}

// GetDevices for a specific group
func (g *Groups) GetDevices(name string) []*Device {
	var devices []*Device
	for _, device := range *g.Devices {
		for _, group := range device.Groups {
			if group == name {
				devices = append(devices, device)
			}
		}
	}

	return devices
}

// NewGroups ...
func NewGroups(devices *Devices) Groups {
	return Groups{
		Devices: devices,
	}
}
