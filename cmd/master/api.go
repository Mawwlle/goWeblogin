package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Routes(e *echo.Echo, app *App) {
	e.GET("/test", func(context echo.Context) error {
		return context.JSON(
			http.StatusCreated, "If you see this message, then all works good")
	})
	e.POST("/adduser", Adder(*app))
	e.POST("/getuser", Getter(*app))
}
