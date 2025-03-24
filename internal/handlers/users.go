package handlers

import (
	"net/http"

	"backend/internal/model"

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
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Ошибка при сохранении пользователя"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "User registered successfully", "token": "token"})
}
