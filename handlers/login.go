package handlers

import (
	"net/http"

	"github.com/abdullahshafaqat/GOTASKS/services"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

var login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LogIn(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		if err := c.ShouldBindJSON(&login); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Data not in json"})
			return
		}
		var user Info

		err := db.Get(&user, `SELECT username, email, password FROM users WHERE email = $1`, login.Email)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password)); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
			return
		}

		access, refresh,err := services.GenerateTokens(user.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"access_token":   access,
			"refresh_token":   refresh,
		})
	}
}
