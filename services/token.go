package services

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var accessKey = []byte(os.Getenv("AC_SECRET"))
var refreshKey = []byte(os.Getenv("RF_SECRET"))

func GenerateTokens(email string) (string, string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"type":  "access",
		"exp":   time.Now().Add(time.Minute * 1).Unix(),
	})

	accessTokenString, err := accessToken.SignedString(accessKey)
	if err != nil {
		return "", "", err
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"type":  "refresh",
		"exp":   time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	refreshTokenString, err := refreshToken.SignedString(refreshKey)
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return accessKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("token expired")
	}

	return nil
}
func VerifyRefreshToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return refreshKey, nil
	})

	if err != nil {
		return "", err
	}

	if !token.Valid{
		return "", fmt.Errorf("invalid refresh token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || claims["email"] == nil {
		return "", fmt.Errorf("invalid claims in token")
	}

	email := claims["email"].(string)
	return email, nil
}
