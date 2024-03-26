package handler

import (
	"net/http"
	"todo/data"
	"todo/view/users"

	"github.com/anthdm/slick"
)

func HandleAuthPage(c *slick.Context) error {
	return c.Render(users.AuthPage())
}

func HandleLoginPost(c *slick.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	// check username and password against database
	user, err := data.UserGet(username)
	if err != nil {
		c.Render(users.AuthPage())
		return err
	}

	// check password
	match, err := data.UserPasswordMatches(password, user)
	if err != nil {
		c.Render(users.AuthPage())
		return err
	}
	if !match {
		c.Render(users.AuthPage())
		return nil
	}

	return c.Redirect("/todos", http.StatusFound)
}
