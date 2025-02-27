package app

import (
	"backend/internal/middleware"
	"backend/internal/routes"
	"log/slog"

	"github.com/labstack/echo/v4"
)

func New(log *slog.Logger, port string) {
	e := echo.New()

	routes.New(e, log)
	middleware.New(e)

	if err := e.Start(":" + port); err != nil {
		log.Error("Failed to start server", slog.Any("error", err))
	}
}
