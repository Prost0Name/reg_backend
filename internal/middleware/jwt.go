package middleware

// import (
// 	"time"

// 	"backend/internal/config"

// 	"github.com/golang-jwt/jwt/v5"
// )

// var jwtKey = []byte(config.JwtSecret)

// type Claims struct {
// 	Login string `json:"login"`
// 	jwt.RegisteredClaims
// }

// func GenerateToken(login string) (string, error) {
// 	expirationTime := time.Now().Add(24 * time.Hour)
// 	claims := &Claims{
// 		Login: login,
// 		RegisteredClaims: jwt.RegisteredClaims{
// 			ExpiresAt: jwt.NewNumericDate(expirationTime),
// 		},
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	return token.SignedString(jwtKey)
// }

// func ValidateToken(tokenString string) (*Claims, error) {
// 	claims := &Claims{}
// 	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
// 		return jwtKey, nil
// 	})

// 	if err != nil || !token.Valid {
// 		return nil, err
// 	}
// 	return claims, nil
// }
