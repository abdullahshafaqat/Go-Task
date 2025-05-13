package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type Info struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SignUp(db *sqlx.DB ) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user Info

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Data not in json"})
			return
		}
		var exists bool
		err := db.QueryRow(
			`SELECT EXISTS(SELECT 1 FROM users WHERE username = $1 OR email = $2)`,
			user.Username, user.Email,
		).Scan(&exists)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Errors of database"})
			return
		}

		if exists {
			c.JSON(http.StatusConflict, gin.H{"error": "Username or email already exists"})
			return
		}
		_, err = db.Exec(
			`INSERT INTO users (username, email, password) VALUES ($1, $2, $3)`,
			user.Username, user.Email, user.Password,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "User created sucessfully",
			"user":    user.Username,
		})
	}
}
