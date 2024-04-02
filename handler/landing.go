package handler

import (
	"net/http"
	"todo/data"
	"todo/view/landing"

	"github.com/anthdm/slick"
)

func HandleLandingPage(c *slick.Context) error {
	user, _ := data.UserGetUser(c)
	if user != nil {
		return c.Redirect("/todos", http.StatusFound)
	}
	return c.Render(landing.LandingPage())
}
