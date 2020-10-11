package main

import (
	"fmt"

	"github.com/br4in3x/google-action-time-since/actions"
	"github.com/br4in3x/google-action-time-since/internal/util"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/tidwall/pretty"
)

func main() {
	e := echo.New()
	e.Logger.SetLevel(log.INFO)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		payload := string(pretty.Color(reqBody, nil))
		fmt.Printf("%s\n", payload)
	}))

	srv := Server{
		actions.NewDateFromToAction(&util.TimeWrapper{}),
		actions.NewDateBetweenAction(&util.TimeWrapper{}),
	}

	e.POST("/webhook", srv.Webhook)

	e.Logger.Fatal(e.Start("0.0.0.0:8080"))
}
