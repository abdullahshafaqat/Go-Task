package handlers

import (
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type Info struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func isValidGmail(email string) bool {
	valid := regexp.MustCompile(`^[^@]+@gmail\.com$`)
	return valid.MatchString(email)
}

func SignUp(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user Info

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON input"})
			return
		}

		if !isValidGmail(user.Email) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Enter a valid Email address"})
			return
		}
		_ = db.QueryRow(
			`SELECT EXISTS(SELECT 1 FROM users WHERE username = $1 OR email = $2)`,
			user.Username, user.Email,
		)
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		_ = db.QueryRow(
			`INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id`,
			user.Username, user.Email, string(hashedPassword),
		).Scan(&user.ID)

		c.JSON(http.StatusCreated, gin.H{
			"message": "User created successfully",
			"user":    user.Username,
		})
	}
}
