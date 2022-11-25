package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

var (
	VendorName      = "knackwurstking"
	ApplicationName = "alice"
	UserShareDir    = filepath.Join(
		userHomeDir(), ".local", "share", VendorName, ApplicationName,
	)
)

type Server struct {
	Name string `json:"name"`
	Port int    `json:"port"`
}

type Servers []*Server

// SetServer to "*/alice/config.json" (panic if error)
func SetServer(server *Server) {
	// This part is maybe a little bit weird!
	//  -> "Add current server to the alice config.json file, if not already exists" - used from talice (telegram bot)
	configPath := filepath.Join(UserShareDir, "config.json")
	data, err := os.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(UserShareDir, 0700)
			if err != nil {
				panic(err)
			}
			data, err = json.Marshal(server)
			if err != nil {
				panic(err)
			}
			err = os.WriteFile(configPath, data, 0644)
		}

		if err != nil {
			panic(err)
		}
	} else {
		var servers Servers
		err := json.Unmarshal(data, &servers)
		if err != nil {
			panic(err)
		}

		for _, s := range servers {
			if server.Name == s.Name && server.Port == s.Port {
			    // server exists
				return
			}
		}

		// merge
		servers = append(servers, server)

		// update
		data, err = json.Marshal(servers)
		if err != nil {
			panic(err)
		}

        // write to file
		err = os.WriteFile(configPath, data, 0644)
		if err != nil {
			panic(err)
		}
	}
}

func userHomeDir() (home string) {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return
}
