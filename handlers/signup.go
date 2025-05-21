package handlers

import (
	"net/http"
	"regexp"

	"github.com/abdullahshafaqat/GOTASKS/api/config/repositories"
	"github.com/abdullahshafaqat/GOTASKS/models"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func isValidGmail(email string) bool {
	valid := regexp.MustCompile(`^[^@]+@gmail\.com$`)
	return valid.MatchString(email)
}

func SignUp(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var User models.NewUser

		if err := c.ShouldBindJSON(&User); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON input"})
			return
		}

		if !isValidGmail(User.Email) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Enter a valid Email address"})
			return
		}
		_ = db.QueryRow(
			`SELECT EXISTS(SELECT 1 FROM users WHERE username = $1 OR email = $2)`,
			User.Username, User.Email,
		)
		hashedPassword, _ := repositories.HashedPassword(User)
		_ = db.QueryRow(
			`INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id`,
			User.Username, User.Email, string(hashedPassword),
		)

		c.JSON(http.StatusCreated, gin.H{
			"message": "User created successfully",
			"user":    User.Username,
		})
	}
}
