package handlers

import (
	"net/http"

	"github.com/abdullahshafaqat/GOTASKS/models"
	"github.com/abdullahshafaqat/GOTASKS/services"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

func LogIn(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var entry models.Credentials
		var info models.NewUser
		if err := c.ShouldBindJSON(&entry); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		err := db.Get(&info, `SELECT id, username, email, password FROM users WHERE email = $1`, entry.Email)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(info.Password), []byte(entry.Password)); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
			return
		}

		access, refresh, _ := services.GenerateTokens(info.ID)
		c.JSON(http.StatusOK, gin.H{
			"access_token":  access,
			"refresh_token": refresh,
		})
	}
}
