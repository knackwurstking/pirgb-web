package middleware

import (
	"crypto/sha256"
	"crypto/subtle"
	"net/http"
	"os"
)

func Auth(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ww := GetCachedResponse(w)

		// check headers for authorization key
		user, pass, ok := r.BasicAuth()
		if checkAuth(user, pass, ok) {
			h.ServeHTTP(ww, r)
		}

		// exit status code: 401 (StatusUnauthorized)
		ww.Header().Set("WWW-Authenticate", `Basic realm="restricted, charset="UTF-8"`)
		http.Error(ww, http.StatusText(http.StatusUnauthorized),
			http.StatusUnauthorized)
	}
}

func checkAuth(user, pass string, ok bool) bool {
	// check user and pass and serve if ok
	userHash := sha256.Sum256([]byte(user))
	passHash := sha256.Sum256([]byte(pass))
	expectedUserHash := sha256.Sum256([]byte(os.Getenv("PIRGB_WEB_USER")))
	expectedPassHash := sha256.Sum256([]byte(os.Getenv("PIRGB_WEB_PASS")))

	userMatch := subtle.ConstantTimeCompare(userHash[:], expectedUserHash[:]) == 1
	passMatch := subtle.ConstantTimeCompare(passHash[:], expectedPassHash[:]) == 1
	if userMatch && passMatch {
		return true
	}

	return false
}
