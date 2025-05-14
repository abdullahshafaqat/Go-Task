package handlers

import (
	"net/http"

	"github.com/abdullahshafaqat/GOTASKS/services"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Authorization(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Content-Type", "application/json")
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			c.String(http.StatusUnauthorized, "Missing authorization header")
			return
		}
		tokenString = tokenString[len("Bearer "):]

		err := services.VerifyToken(tokenString)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			c.String(http.StatusUnauthorized, "Invalid token")
			return
		}

		c.String(http.StatusOK, "You are autherized")

	}
}
