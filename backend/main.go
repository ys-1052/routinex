// Package main provides the entry point for the RoutineX backend server.
package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "RoutineX API",
		})
	})

	if err := e.Start(":8080"); err != nil {
		e.Logger.Fatal("Server failed to start:", err)
	}
}
