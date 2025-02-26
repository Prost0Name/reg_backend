package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"backend/internal/middleware"
	"backend/internal/models"
	"backend/internal/types"
)

func Register(c echo.Context) error {
	creds := new(types.Credentials)
	if err := c.Bind(creds); err != nil {
		return c.String(http.StatusBadRequest, "Неверный формат данных")
	}

	exists, err := models.UserExists(creds.Login)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Ошибка при проверке пользователя")
	}

	if exists {
		return c.JSON(http.StatusConflict, map[string]string{
			"message": "Пользователь с таким логином уже существует",
		})
	}

	if err := models.AddUser(creds.Login, creds.Password); err != nil {
		return c.String(http.StatusInternalServerError, "Ошибка при регистрации пользователя")
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Регистрация успешно завершена",
	})
}

func Auth(c echo.Context) error {
	creds := new(types.Credentials)
	if err := c.Bind(creds); err != nil {
		return c.String(http.StatusBadRequest, "Неверный формат данных")
	}

	// Проверяем учетные данные пользователя
	valid, err := models.ValidateUser(creds.Login, creds.Password)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Ошибка при проверке учетных данных")
	}

	if !valid {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "Неверный логин или пароль",
		})
	}

	// Генерация токена
	token, err := middleware.GenerateToken(creds.Login)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Ошибка при генерации токена")
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Аутентификация успешна",
		"token":   token,
	})
}
