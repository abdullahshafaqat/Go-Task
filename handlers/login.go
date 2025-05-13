package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
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

		err := db.Get(&user, `SELECT username,email,password
			FROM users
			 WHERE email = $1 AND  password = $2`,
			login.Email, login.Password)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Login successfully",
		})
	}
}
