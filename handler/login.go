package handler

import (
	"net/http"
	"todo/configs"
	"todo/data"
	"todo/session"
	"todo/types"
	"todo/view/auth"

	"github.com/anthdm/slick"
)

func HandleLoginPage(c *slick.Context) error {
	return c.Render(auth.LoginPage())
}

func HandleLoginPost(c *slick.Context) error {
	session, err := configs.Store.Get(c.Request, "user-session")
	if err != nil {
		http.Error(c.Response, err.Error(), http.StatusInternalServerError)
		return err
	}

	username := c.FormValue("username")
	password := c.FormValue("password")
	// check username and password against database
	user, err := data.UserGet(username)
	if err != nil {
		c.Render(auth.LoginPage())
		return err
	}

	// check password
	match, err := data.UserPasswordMatches(password, user)
	if err != nil {
		c.Render(auth.LoginPage())
		return err
	}
	if !match {
		c.Render(auth.LoginPage())
		return nil
	}

	authenticatedUser := &types.AuthenticatedUser{
		UUID:          user.UUID,
		Username:      user.Username,
		Email:         user.Email,
		Authenticated: true,
	}

	// store user id in session
	session.Values["user"] = authenticatedUser

	err = session.Save(c.Request, c.Response)
	if err != nil {
		http.Error(c.Response, err.Error(), http.StatusInternalServerError)
		return err
	}

	return c.Redirect("/todos", http.StatusFound)
}

func HandleLogoutGet(c *slick.Context) error {
	userSession, err := configs.Store.Get(c.Request, "user-session")
	if err != nil {
		http.Error(c.Response, err.Error(), http.StatusInternalServerError)
		return err
	}
	// remove user from session TODO: refactor to sessison
	userSession.Values["user"] = nil
	err = userSession.Save(c.Request, c.Response)
	if err != nil {
		http.Error(c.Response, err.Error(), http.StatusInternalServerError)
		return err
	}

	session.AddAlert(c, &types.AlertType{Type: types.AlertEnum.Success, Message: "Logged out"})

	return c.Redirect("/login", http.StatusFound)
}
