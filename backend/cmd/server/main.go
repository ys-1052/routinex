// Package main provides the main server entry point for RoutineX backend.
package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/ys-1052/routinex/backend/internal/handlers"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "RoutineX API",
		})
	})

	e.GET("/health", handlers.Health)

	if err := e.Start(":8080"); err != nil {
		e.Logger.Fatal("Server failed to start:", err)
	}
}
