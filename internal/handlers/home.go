package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Home(c echo.Context) error {
	// Проверяем наличие токена в контексте
	login := c.Get("login")
	if login == nil {
		// Если токен отсутствует, перенаправляем на главную страницу
		return c.Redirect(http.StatusFound, "/") // Замените "/" на URL вашей главной страницы
	}

	// Если токен действителен, возвращаем страницу /home
	return c.String(http.StatusOK, "Добро пожаловать на страницу /home, "+login.(string))
}
