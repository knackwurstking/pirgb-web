package main

import (
	"encoding/json"
	"fmt"

	"gitlab.com/knackwurstking/pirgb-web/internal/config"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	data, _ := json.Marshal(config.GlobalData)
	fmt.Printf("%s\n", data)
}
