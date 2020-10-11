package main

import (
	"fmt"
	"os"

	"github.com/br4in3x/time-since-google-action/actions"
	"github.com/br4in3x/time-since-google-action/internal/util"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tidwall/pretty"
)

func main() {
	e := echo.New()

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

	port := os.Getenv("PORT")
	addr := fmt.Sprintf("0.0.0.0:%s", port)
	e.Logger.Fatal(e.Start(addr))
}
