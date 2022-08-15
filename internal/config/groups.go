package config

// Groups ...
type Groups struct {
	Devices *Devices `json:"-"`
}

// GetGroups for all devices
func (g *Groups) GetGroups() []string {
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

// TODO: list groups available: []string
// TODO: get group: returns a new filtered list of devices []*Devices

// NewGroups ...
func NewGroups(devices *Devices) Groups {
	return Groups{
		Devices: devices,
	}
}
