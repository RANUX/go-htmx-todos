package middleware

import (
	"net/http"
	"todo/data"
	"todo/pkg/session"

	"github.com/anthdm/slick"
)

// AuthMiddleware is middleware for user authentication
func AuthMiddleware(next slick.Handler) slick.Handler {
	return func(c *slick.Context) error {
		// Before handler logic here, e.g., logging, authentication
		r := c.Request
		w := c.Response
		userSession, err := session.Store.Get(r, "user-session")

		if err != nil {
			http.Redirect(w, r, "/login", http.StatusFound) // TODO: redirect to logout
			return err
		}

		authUser := session.GetAuthenticatedUser(userSession)
		if !authUser.Authenticated {
			http.Redirect(w, r, "/login", http.StatusFound)
			return err
		}

		user, err := data.UserGetByUUID(authUser.UUID)
		if err != nil || user == nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return err
		}

		// Run next handler
		return next(c)
	}

	// If user is not authenticated, redirect to login page
	//h(w, r, ps)
}
