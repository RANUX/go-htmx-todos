package session

import (
	"todo/types"

	"github.com/anthdm/slick"

	"github.com/sirupsen/logrus"
)

func PopAlerts(c *slick.Context) []*types.AlertType {
	s, err := Store.Get(c.Request, "alert-session")
	alerts := make([]*types.AlertType, 0)

	if err != nil {
		logrus.Error(err)
		return alerts
	}
	val := s.Values["alerts"]

	alerts, ok := val.([]*types.AlertType)
	if !ok {
		return make([]*types.AlertType, 0)
	}

	// clear alerts
	s.Values["alerts"] = nil
	err = s.Save(c.Request, c.Response)
	if err != nil {
		logrus.Error(err)
		return make([]*types.AlertType, 0)
	}

	return alerts
}

func AppendAlert(c *slick.Context, alert *types.AlertType) {
	s, err := Store.Get(c.Request, "alert-session")
	if err != nil {
		logrus.Error(err)
		return
	}
	alerts := PopAlerts(c)
	alerts = append(alerts, alert)
	s.Values["alerts"] = alerts
	err = s.Save(c.Request, c.Response)
	if err != nil {
		logrus.Error(err)
		return
	}
}
