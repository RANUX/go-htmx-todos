package plug

import (
	"net/http"
	"todo/configs"
	"todo/data"
	"todo/session"
	"todo/types"

	"github.com/anthdm/slick"
)

// AuthMiddleware is middleware for user authentication
func AuthMiddleware(next slick.Handler) slick.Handler {
	return func(c *slick.Context) error {
		// Before handler logic here, e.g., logging, authentication
		r := c.Request
		w := c.Response
		userSession, err := configs.Store.Get(r, "user-session")

		if err != nil {
			session.AddAlert(c, &types.AlertType{Type: types.AlertEnum.Info, Message: "Please login to continue"})
			http.Redirect(w, r, "/login", http.StatusFound)
			return err
		}

		authUser := session.GetAuthenticatedUser(userSession)
		if !authUser.Authenticated {
			session.AddAlert(c, &types.AlertType{Type: types.AlertEnum.Info, Message: "Please login to continue"})
			http.Redirect(w, r, "/login", http.StatusFound)
			return err
		}

		user, err := data.UserGetByUUID(authUser.UUID)
		if err != nil || user == nil {
			session.AddAlert(c, &types.AlertType{Type: types.AlertEnum.Info, Message: "Please login to continue"})
			http.Redirect(w, r, "/login", http.StatusFound)
			return err
		}

		c.Set(string(types.ContextEnum.User), user)

		// Run next handler
		return next(c)
	}

}
