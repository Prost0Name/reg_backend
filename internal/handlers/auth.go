package handlers

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"

	"backend/internal/config"
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

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"login": creds.Login,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString(config.JwtSecret)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Ошибка при создании токена")
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Аутентификация успешна",
		"token":   tokenString,
	})
}
