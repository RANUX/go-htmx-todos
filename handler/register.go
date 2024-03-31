package handler

import (
	"todo/view/auth"

	"github.com/anthdm/slick"
)

func HandleRegisterPage(c *slick.Context) error {
	return c.Render(auth.RegisterPage())
}
