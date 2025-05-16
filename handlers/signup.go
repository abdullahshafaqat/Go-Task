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
		err := db.QueryRow(
			`SELECT EXISTS(SELECT 1 FROM users WHERE username = $1 OR email = $2)`,
			user.Username, user.Email,
		).Scan(&exists)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database check error"})
			return
		}

		if exists {
			c.JSON(http.StatusConflict, gin.H{"error": "Username or email already exists"})
			return
		}
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Password hashing failed"})
			return
		}

		_, err = db.Exec(
			`INSERT INTO users (username, email, password) VALUES ($1, $2, $3)`,
			user.Username, user.Email, string(hashedPassword),
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "User creation failed"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "User created successfully",
			"user":    user.Username,
		})
	}
}
