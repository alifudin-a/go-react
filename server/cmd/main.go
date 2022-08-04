package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "==> STATUS=${status}, METHOD=${method}, HOST=${host}, URI=${uri}, " +
			"ERROR=${error}, LATENCY_HUMAN=${latency_human}\n",
	}))

	e.Static("/", "../client/build")

	e.Logger.Fatal(e.Start(":5000"))
}
