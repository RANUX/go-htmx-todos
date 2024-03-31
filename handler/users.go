package handler

import (
	"net/http"
	"todo/data"
	"todo/pkg/session"
	"todo/types"
	"todo/view/users"

	"github.com/anthdm/slick"
)

func HandleProfilePage(c *slick.Context) error {
	userSession, err := session.Store.Get(c.Request, "user-session")
	if err != nil {
		http.Error(c.Response, err.Error(), http.StatusInternalServerError)
		return err
	}

	authUser := session.GetAuthenticatedUser(userSession)

	// get user from db by username
	user, err := data.UserGetByUUID(authUser.UUID)
	if err != nil {
		return err
	}

	return c.Render(users.ProfilePage(users.ProfileProps{User: *user}))
}

func HandleProfileSave(c *slick.Context) error {
	userSession, err := session.Store.Get(c.Request, "user-session")
	if err != nil {
		http.Error(c.Response, err.Error(), http.StatusInternalServerError)
		return err
	}

	username := c.FormValue("username")
	password := c.FormValue("password")

	authUser := session.GetAuthenticatedUser(userSession)

	// get user from db by username
	user, err := data.UserGetByUUID(authUser.UUID)
	if err != nil {
		return err
	}

	// update username
	user.Username = username

	if password != "" {
		newHash, err := data.HashPassword(password)
		if err != nil {
			return err
		}
		user.Password = newHash
	}

	err = data.UserUpdate(user)
	if err != nil {
		return err
	}

	session.AddAlert(c, &types.AlertType{Type: types.AlertEnum.Success, Message: "Profile updated successfully"})

	return c.Render(users.ProfileForm(users.ProfileProps{User: *user}))
}
