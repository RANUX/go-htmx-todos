package session

import (
	"net/http"
	"todo/configs"
	"todo/types"

	"github.com/anthdm/slick"
)

func GetAuthenticatedUser(c *slick.Context) (types.AuthenticatedUser, error) {
	s, err := configs.Store.Get(c.Request, "user-session")
	if err != nil {
		http.Error(c.Response, err.Error(), http.StatusInternalServerError)
		return types.AuthenticatedUser{Authenticated: false}, err
	}

	val := s.Values["user"]
	var user = types.AuthenticatedUser{}
	user, ok := val.(types.AuthenticatedUser)
	if !ok {
		return types.AuthenticatedUser{Authenticated: false}, nil
	}
	return user, nil
}
