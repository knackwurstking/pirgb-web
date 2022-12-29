package middleware

import (
	"crypto/sha256"
	"crypto/subtle"
	"net/http"

	"github.com/knackwurstking/pirgb-web/pkg/log"
)

type User struct {
	Name     string
	Password string
}

func (user *User) NameHash() [32]byte {
	return sha256.Sum256([]byte(user.Name))
}

func (user *User) PasswordHash() [32]byte {
	return sha256.Sum256([]byte(user.Password))
}

func (user *User) Verify(name, password string) bool {
	nameHash := sha256.Sum256([]byte(name))
	passwordHash := sha256.Sum256([]byte(password))
	expectedNameHash := user.NameHash()
	expectedPasswordHash := user.PasswordHash()

	nameMatch := subtle.ConstantTimeCompare(nameHash[:], expectedNameHash[:]) == 1
	passwordMatch := subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:]) == 1

	if nameMatch && passwordMatch {
		return true
	}

	return false
}

// Authorization middleware
func Authorization(h http.HandlerFunc, user *User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ww := GetCachedResponse(w)

		log.Debug.Printf("PROTECTED ROUTE: %s", r.URL.Path)

		username, password, ok := r.BasicAuth()
		if ok {
			if user.Verify(username, password) {
				h.ServeHTTP(ww, r)
				return
			}
		}

		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
	}
}
