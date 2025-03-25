package handlers

import (
	"net/http"
	"time"

	"backend/internal/model"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type RegisterRequest struct {
	Login    string `json:"login" form:"login" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

func Register(c echo.Context) error {
	var req RegisterRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	if req.Login == "" || req.Password == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Поле не может быть пустым"})
	}

	if err := model.CreateUser(req.Login, req.Password); err != nil {
		if err.Error() == "пользователь уже существует" {
			return c.JSON(http.StatusConflict, map[string]string{"error": "Пользователь уже существует"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Ошибка при сохранении пользователя"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "User registered successfully"})
}

type LoginRequest struct {
	Login    string `json:"login" form:"login" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

func Login(c echo.Context) error {
	var req LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	user, err := model.GetUserByLogin(req.Login)
	if err != nil || user.Password != req.Password {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid login or password"})
	}

	token, err := generateJWT(req.Login)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Ошибка при создании токена"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Login successful", "token": token})
}

func generateJWT(login string) (string, error) {
	claims := &jwt.MapClaims{
		"login": login,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("your_secret_key")) // Replace "your_secret_key" with your actual secret
}
