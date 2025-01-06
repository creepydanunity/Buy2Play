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
	IsAdmin  bool   `json:"isAdmin"`
	jwt.StandardClaims
}

func GenerateToken(userID uint, username, email string, isAdmin bool, hours int) (string, error) {
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
		isAdmin,
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

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, err
	}

	if claims.ExpiresAt < time.Now().Unix() {
		return nil, jwt.ErrSignatureInvalid
	}

	return claims, nil
}
