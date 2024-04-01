package handler

import (
	"net/http"
	"todo/session"
	"todo/view/alert"

	"github.com/anthdm/slick"
)

func HandleAlerts(c *slick.Context) error {
	alerts := session.PopAlerts(c)

	if len(alerts) == 0 {
		return c.Text(http.StatusOK, "")
	}

	props := alert.AlertListProps{
		Alerts: alerts,
	}
	return c.Render(alert.AlertList(props))
}
