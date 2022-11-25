package api

import (
	"net/http"

	aliceApiMiddleware "github.com/knackwurstking/alice/api/v1/middleware"
	aliceApiTypes "github.com/knackwurstking/alice/api/v1/types"
	"github.com/knackwurstking/alice/pkg/auth"
	"github.com/knackwurstking/pirgb-web/internal/constants"
)

type Auth struct{}

func Authorization(w *aliceApiTypes.ResponseWriter, r *http.Request) (*auth.Authorization, error) {
	return aliceApiMiddleware.Authorization(constants.ApplicationName, constants.AuthDB, w, r)
}
