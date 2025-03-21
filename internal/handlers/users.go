package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type RegisterRequest struct {
	Login    string `json:"login" form:"login" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type User struct {
	Login    string
	Password string
}

func (req *RegisterRequest) ToUser() User {
	return User{
		Login:    req.Login,
		Password: req.Password,
	}
}

func Register(c echo.Context) error {
	var req RegisterRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	user := req.ToUser()

	if user.Login == "" || user.Password == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Поле не может быть пустым"})
	}

	// if err := models.AddUser(user.Login, user.Password); err != nil {
	// 	return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	// }

	return c.JSON(http.StatusOK, map[string]string{"message": "User registered successfully", "token": "token"})
}
