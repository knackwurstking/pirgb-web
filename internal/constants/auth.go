package constants

import (
	"github.com/knackwurstking/alice/pkg/auth"
	aliceConfig "github.com/knackwurstking/alice/pkg/config"
	"github.com/knackwurstking/pirgb-web/pkg/log"
)

var AuthDB *auth.DB

func init() {
	defer func() {
		if r := recover(); r != nil {
			log.Warn.Println(r)
		}
	}()

	AuthDB = auth.NewDB(aliceConfig.AuthDBPath)
}
