package handlers

import (
	"net/http"

	"github.com/abdullahshafaqat/GOTASKS/models"
	"github.com/abdullahshafaqat/GOTASKS/services"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func LogIn(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var entry models.LoginUser
		if err := c.ShouldBindJSON(&entry); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}
		access, refresh, err := services.LoginUser(db, entry)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"access_token":  access,
			"refresh_token": refresh,
		})
	}
}
