package main

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/medeirosvictor/routiny/dto"
	"github.com/medeirosvictor/routiny/templates"
	"github.com/medeirosvictor/routiny/templates/components"
)

func main() {
	e := echo.New()
	e.Static("/static", "static")

	// Mock tasks, fetch from firebase later perhaps
	tasks := []*dto.TaskDto{
		{
			ID:   "1",
			Text: "Draw",
			Done: false,
		},
		{
			ID:   "2",
			Text: "Dev",
			Done: false,
		},
	}

	e.GET("/", func(c echo.Context) error {
		component := templates.Home(tasks)
		return component.Render(c.Request().Context(), c.Response().Writer)
	})

	e.POST("/add-task", func(c echo.Context) error {
		task := dto.TaskDto{ID: uuid.New().String(), Text: c.FormValue("task"), Done: false}
		tasks = append(tasks, &task)
		component := components.TaskTile(task)
		return component.Render(c.Request().Context(), c.Response().Writer)
	})

	e.POST("/remove-task", func(c echo.Context) error {
		task := c.QueryParam("id")
		for i, t := range tasks {
			if t.ID == task {
				tasks = append(tasks[:i], tasks[i+1:]...)
				break
			}
		}
		return c.String(200, "")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
