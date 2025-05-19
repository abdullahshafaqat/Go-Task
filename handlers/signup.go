package handlers

import (
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type Info struct {
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
		var exists bool
		_ = db.QueryRow(
			`SELECT EXISTS(SELECT 1 FROM users WHERE username = $1 OR email = $2)`,
			user.Username, user.Email,
		).Scan(&exists)
		if exists {
			c.JSON(http.StatusConflict, gin.H{"error": "Username or email already exists"})
			return
		}
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		
		_, _ = db.Exec(
			`INSERT INTO users (username, email, password) VALUES ($1, $2, $3)`,
			user.Username, user.Email, string(hashedPassword),
		)

		c.JSON(http.StatusCreated, gin.H{
			"message": "User created successfully",
			"user":    user.Username,
		})
	}
}
