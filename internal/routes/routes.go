package routes

import (
	"backend/internal/handlers"

	"github.com/labstack/echo/v4"

	"log/slog"
)

func New(e *echo.Echo, log *slog.Logger) {
	// e.POST("/auth", handlers.Auth)
	e.POST("/reg", handlers.Register)
}
