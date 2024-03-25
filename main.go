package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/medeirosvictor/routiny/templates"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		component := templates.Home()
		return component.Render(c.Request().Context(), c.Response().Writer)
	})

	e.GET("/add-task", func(c echo.Context) error {
		task := c.QueryParam("task")
		fmt.Println(task)
		return c.String(200, "Task added")
	})

	e.GET("/remove-task", func(c echo.Context) error {
		task := c.QueryParam("task")
		fmt.Println(task)
		return c.String(200, "Task removed")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
