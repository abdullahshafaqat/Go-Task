package handlers

import (
	"net/http"
	"strings"

	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jmoiron/sqlx"
)

var accessKey = []byte(os.Getenv("AC_SECRET"))
var refreshKey = []byte(os.Getenv("RF_SECRET"))

func Authorization(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
			tokenString := c.GetHeader("Authorization")
			if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
				c.String(http.StatusUnauthorized, "Missing authorization header")
				return
			}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return accessKey, nil
		})

		if err == nil && token.Valid {
			if claims, ok := token.Claims.(jwt.MapClaims); ok && claims["type"] == "access" {
				c.String(http.StatusOK, "You are authorized ")
				return
			}
		}

		token, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return refreshKey, nil
		})

		if err == nil && token.Valid {
			if claims, ok := token.Claims.(jwt.MapClaims); ok && claims["type"] == "refresh" {
				c.String(http.StatusOK, "This is a refresh token")
				return
			}
		}

		c.String(http.StatusUnauthorized, "Invalid token")
	}
}
