package middleware

import (
	"backend/internal/config"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)



func jwt(e *echo.Echo) {
	// JWT middleware
	e.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: config.JwtSecret,
		Skipper: func(c echo.Context) bool {
			// Skip authentication for login and register routes
			if c.Path() == "/auth" || c.Path() == "/register" {
				return true
			}
			return false
		},
	}))
}
