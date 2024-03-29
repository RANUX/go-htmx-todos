package session

import (
	"todo/types"

	"github.com/gorilla/sessions"
)

func GetAuthenticatedUser(s *sessions.Session) types.AuthenticatedUser {
	val := s.Values["user"]
	var user = types.AuthenticatedUser{}
	user, ok := val.(types.AuthenticatedUser)
	if !ok {
		return types.AuthenticatedUser{Authenticated: false}
	}
	return user
}
