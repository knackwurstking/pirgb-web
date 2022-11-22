package constants

type Config struct {
	Port    int     `json:"port"`
	Host    string  `json:"host"`
	Devices Devices `json:"devices"`
}

func NewDefaultConfig() *Config {
	return &Config{
		Port:    50831,
		Host:    "",
		Devices: make(Devices, 0),
	}
}

type Devices []Device

type Device struct {
	Host   string   `json:"host"`
	Port   int      `json:"port"`
	Groups []string `json:"groups"`
}
