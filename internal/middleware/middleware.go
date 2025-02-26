package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func New(e *echo.Echo) {
	cors(e)

	// Middleware для проверки JWT токена
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Пропускаем проверку токена для маршрутов регистрации и аутентификации
			if c.Request().URL.Path == "/register" || c.Request().URL.Path == "/auth" {
				return next(c)
			}

			token := c.Request().Header.Get("Authorization")
			if token == "" {
				return c.String(http.StatusUnauthorized, "Токен не предоставлен")
			}

			claims, err := ValidateToken(token)
			if err != nil {
				return c.String(http.StatusUnauthorized, "Неверный токен")
			}

			c.Set("login", claims.Login)
			return next(c)
		}
	})
}
