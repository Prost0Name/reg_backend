package routes

import (
	"backend/internal/handlers"

	"github.com/labstack/echo/v4"
)

func New(e *echo.Echo) {
	e.POST("/auth", handlers.Auth)
	e.POST("/register", handlers.Register)
	e.GET("/home", handlers.Home)
}
