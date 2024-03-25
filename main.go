package main

import (
	"github.com/labstack/echo/v4"
	"github.com/medeirosvictor/routiny/templates"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		component := templates.Home()
		return component.Render(c.Request().Context(), c.Response().Writer)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
