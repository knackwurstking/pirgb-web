package api

import "net/http"

type Auth struct {}

func Authorization(w ResponseWriter, r *http.Request) *Auth {
    // TODO: check authorization (telegram - "talice")

    return nil
}
