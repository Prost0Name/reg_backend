package routes

import (
	"backend/internal/handlers"

	"github.com/labstack/echo/v4"
)

func Users(e *echo.Echo) {
	e.POST("/reg", handlers.Register)
	// e.GET()
}
