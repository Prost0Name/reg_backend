package middleware

import (
	"github.com/labstack/echo/v4"
)

func New(e *echo.Echo) {
	cors(e)
	jwt(e)
}
