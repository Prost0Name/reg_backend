package handlers

import (
	"net/http"
	"time"

	"backend/internal/model"
	"backend/internal/redis"
	"backend/utils"

	"crypto/rand"
	"encoding/base64"
	"encoding/json"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type RegisterRequest struct {
	Login    string `json:"login" form:"login" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
}

func generateRandomToken() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func Register(c echo.Context) error {
	var req RegisterRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	if req.Login == "" || req.Password == "" || req.Email == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Все поля обязательны для заполнения"})
	}

	token, err := generateRandomToken()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Ошибка при генерации токена"})
	}

	// Store user data in Redis
	userData := map[string]string{
		"login":    req.Login,
		"email":    req.Email,
		"password": req.Password,
	}
	if err := redis.SetUserData(token, userData); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Ошибка при сохранении данных в Redis"})
	}

	// Send confirmation email with the token link
	subject := "Подтверждение регистрации"
	body := "Пожалуйста, подтвердите вашу регистрацию, перейдя по следующей ссылке: http://yourdomain.com/confirm?token=" + token
	if err := utils.SendEmail(req.Email, subject, body); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Ошибка при отправке письма"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "User registered successfully, please check your email to confirm registration"})
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

func ConfirmRegistration(c echo.Context) error {
	token := c.QueryParam("token")
	userData, err := redis.GetUserData(token)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid or expired token"})
	}

	// Parse userData and create user in DB
	var user map[string]string
	if err := json.Unmarshal([]byte(userData), &user); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Ошибка при обработке данных пользователя"})
	}

	if err := model.CreateUser(user["login"], user["email"], user["password"]); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Ошибка при сохранении пользователя в БД"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Registration confirmed successfully"})
}
