package session

import (
	"net/http"
	"todo/configs"
	"todo/types"

	"github.com/anthdm/slick"
	"github.com/sirupsen/logrus"
)

func PopAlerts(c *slick.Context) []*types.AlertType {
	s, err := configs.Store.Get(c.Request, "flash-session")
	alerts := make([]*types.AlertType, 0)

	if err != nil {
		logrus.Error(err)
		return alerts
	}

	flashMessages := s.Flashes("alert")

	for _, msg := range flashMessages {
		if alert, ok := msg.(types.AlertType); ok {
			alerts = append(alerts, &alert)
		} else {
			// Log the error or handle it as necessary
			logrus.Error("Type assertion to *types.AlertType failed")
		}
	}

	s.Save(c.Request, c.Response) // Consider handling the error from Save
	return alerts
}

func AddAlert(c *slick.Context, alert *types.AlertType) {
	s, err := configs.Store.Get(c.Request, "flash-session")
	if err != nil {
		http.Error(c.Response, err.Error(), http.StatusInternalServerError)
		return
	}

	// Add the single alert as a flash message.
	s.AddFlash(alert, "alert")

	err = s.Save(c.Request, c.Response)
	if err != nil {
		http.Error(c.Response, err.Error(), http.StatusInternalServerError)
		return
	}
}
