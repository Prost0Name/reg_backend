package middleware

import (
	"backend/internal/config"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func jwtMiddleware(e *echo.Echo) {
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

	// Middleware to check for valid JWT token
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cookie, err := c.Cookie("jwt")
			if err != nil {
				// Redirect to homepage if cookie is not found
				return c.Redirect(http.StatusFound, "/")
			}

			// Validate the token
			token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
				// Validate the algorithm
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, echo.NewHTTPError(http.StatusUnauthorized, "Invalid signing method")
				}
				return config.JwtSecret, nil
			})

			if err != nil || !token.Valid {
				// Redirect to homepage if token is invalid
				return c.Redirect(http.StatusFound, "/")
			}

			return next(c)
		}
	})
}
