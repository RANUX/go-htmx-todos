package handler

import (
	"todo/view/landing"

	"github.com/anthdm/slick"
)

func HandleLandingPage(c *slick.Context) error {
	return c.Render(landing.LandingPage())
}
