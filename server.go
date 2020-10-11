package main

import (
	"net/http"

	"github.com/br4in3x/google-action-time-since/actions"
	"github.com/br4in3x/google-action-time-since/internal/request"
	"github.com/br4in3x/google-action-time-since/internal/response"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type ActionFunc = func(r request.Request) (*response.Simple, error)

const (
	ActionDateFrom    = "DATE_FROM_TO"
	ActionDateBetween = "DATE_BETWEEN"
)

type Server struct {
	DateFromAction    *actions.DateFromToAction
	DateBetweenAction *actions.DateBetweenAction
}

func (s Server) DetermineAction(action string) (ActionFunc, error) {
	switch action {
	case ActionDateFrom:
		return s.DateFromAction.Invoke, nil
	case ActionDateBetween:
		return s.DateBetweenAction.Invoke, nil
	default:
		return nil, errors.New("can not determine action")
	}
}

func (s Server) Webhook(c echo.Context) error {
	req := request.Request{}
	if err := c.Bind(&req); err != nil {
		return errors.New("can not unmarshall request payload")
	}

	act, err := s.DetermineAction(req.Handler.Name)
	if err != nil {
		return errors.Wrap(err, "main srv.DetermineAction error")
	}

	res, err := act(req)
	if err != nil {
		return errors.Wrap(err, "main act(req) error")
	}

	return c.JSON(http.StatusOK, res)
}
