// Package handlers provides HTTP request handlers for the RoutineX API.
package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Health handles the health check endpoint and returns the service status.
func Health(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "healthy",
		"message": "RoutineX API is running",
		"service": "routinex-backend",
	})
}
