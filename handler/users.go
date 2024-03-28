package handler

import (
	"todo/data"
	"todo/view/users"

	"github.com/anthdm/slick"
)

func HandleProfilePage(c *slick.Context) error {
	return c.Render(users.ProfilePage())
}

func HandleProfileSave(c *slick.Context) error {

	username := c.FormValue("username")
	password := c.FormValue("password")

	// get user from db by username
	user, err := data.UserGet(username)
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

	return c.Render(users.ProfilePage())
}
