package handler

import (
	"todo/data"
	"todo/session"
	"todo/types"
	"todo/view/users"

	"github.com/anthdm/slick"
)

func HandleProfilePage(c *slick.Context) error {
	// get user from db by username
	user, err := data.UserGetUser(c)
	if err != nil {
		return err
	}

	return c.Render(users.ProfilePage(users.ProfileProps{User: *user}))
}

func HandleProfileSave(c *slick.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	// get user from db by username
	user, err := data.UserGetUser(c)
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
