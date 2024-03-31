package plug

import (
	"net/http"
	"todo/data"
	"todo/pkg/session"
	"todo/types"

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
			http.Redirect(w, r, "/login", http.StatusFound)
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

		c.Set(string(types.ContextEnum.User), user)

		// Run next handler
		return next(c)
	}

}
