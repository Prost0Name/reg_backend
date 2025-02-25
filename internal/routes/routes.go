package routes

import (
	"backend/internal/handlers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func New(e *echo.Echo) {
	e.POST("/auth", handlers.Auth)
	e.POST("/register", handlers.Register)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to the homepage!")
	})
}
