package constants

import (
	"github.com/knackwurstking/alice/pkg/auth"

	aliceConfig "github.com/knackwurstking/alice/pkg/config"
)

var (
    AuthDB *auth.DB = auth.NewDB(aliceConfig.AuthDBPath)
)
