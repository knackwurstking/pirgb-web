package main

import (
	"encoding/json"
	"fmt"

	"gitlab.com/knackwurstking/pirgb-web/internal/config"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)

	// NOTE: just testing ...
	data, _ := json.Marshal(config.GlobalData)
	fmt.Printf("%s\n", data)

	// TODO: parse flags here ...
}
