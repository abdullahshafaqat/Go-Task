package handlers

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jmoiron/sqlx"
)

var accessKey = []byte(os.Getenv("AC_SECRET"))
var refreshKey = []byte(os.Getenv("RF_SECRET"))

func BearerToken(header string) string {
	if strings.HasPrefix(header, "Bearer ") {
		return strings.TrimPrefix(header, "Bearer ")
	}
	return ""
}

func parseToken(tokenString string, secret []byte) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
}

func TokenType(token *jwt.Token, Type string) bool {
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if tokenType, ok := claims["type"].(string); ok && tokenType == Type {
			return true
		}
	}
	return false
}

func Authorization(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := BearerToken(c.GetHeader("Authorization"))
		if tokenString == "" {
			c.String(http.StatusUnauthorized, "Missing or invalid authorization header")
			c.Abort()
			return
		}
		if token, err := parseToken(tokenString, accessKey); err == nil {
			if TokenType(token, "access") {
				
				c.Next()
				return
			}
		}

		if token, err := parseToken(tokenString, refreshKey); err == nil {
			if TokenType(token, "refresh") {
				c.String(http.StatusOK, "This is a refresh token")
				c.Abort()
				return
			}
		} else if strings.Contains(err.Error(), "expired") {
			c.String(http.StatusUnauthorized, "Token is expired")
			c.Abort()
			return
		}

		c.String(http.StatusUnauthorized, "Invalid token")
		c.Abort()
	}
}
