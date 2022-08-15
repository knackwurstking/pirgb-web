package config

const (
	// DefaultPort the pirgb-server default port
	DefaultPort = 50826
)

var (
	// Global contains parse configuration data
	Global = &Config{
		Groups:     make(Groups, 0),
		Devices:    make(Devices, 0),
		EnableHTTP: true,
	}
)
