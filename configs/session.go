package configs

import (
	"encoding/gob"
	"os"

	"todo/types"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/sirupsen/logrus"
)

// store will hold all session data
var Store *sessions.CookieStore

// Initializes the session
func init() {
	// read cookieKey from environment and convert to byte[]
	cookieHashKey := []byte(os.Getenv("COOKIE_HASH_KEY"))

	if len(cookieHashKey) == 0 {
		logrus.Warn("No COOKIE_HASH_KEY key found in env. Using generated key.")
		cookieHashKey = securecookie.GenerateRandomKey(64)
	}

	cookieEncrypted := os.Getenv("COOKIE_ENCRYPTED")
	encryptCookies := true

	if cookieEncrypted == "" || cookieEncrypted == "false" {
		encryptCookies = false
	}

	if encryptCookies {
		cookieBlockKey := []byte(os.Getenv("COOKIE_BLOCK_KEY"))
		if len(cookieBlockKey) == 0 {
			logrus.Warn("No COOKIE_BLOCK_KEY key found in env. Using generated key.")
			cookieBlockKey = securecookie.GenerateRandomKey(32)
		}
		Store = sessions.NewCookieStore(cookieHashKey, cookieBlockKey)
	} else {
		Store = sessions.NewCookieStore(cookieHashKey)
	}

	Store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7, // 7 days
		HttpOnly: true,
		// Secure:   true,
		// SameSite: http.SameSiteLaxMode, // Lax mode to allow CSRF
	}

	gob.Register(types.AuthenticatedUser{})
	gob.Register(types.AlertType{})

}
