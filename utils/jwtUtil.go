package utils

import (
	"buy2play/config"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

func GenerateToken(userID uint, username, email string, hours int) (string, error) {
	var expiresAt int64

	if hours > 10 {
		expiresAt = time.Now().Add(time.Hour * 24 * 14).Unix()
	} else {
		expiresAt = time.Now().Add(time.Hour * 4).Unix()
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		userID,
		username,
		email,
		jwt.StandardClaims{
			ExpiresAt: expiresAt,
			Issuer:    "Buy2Play",
			IssuedAt:  time.Now().Unix(),
		},
	})

	return token.SignedString(config.JWTSecret)
}

func ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return config.JWTSecret, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
