package handlers

import (
	"net/http"

	"github.com/abdullahshafaqat/GOTASKS/models"
	"github.com/abdullahshafaqat/GOTASKS/services"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func SignUp(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var User models.NewUser

		if err := c.ShouldBindJSON(&User); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := services.RegisterUser(db, User)
		if err != nil {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{
			"message": "User created successfully",
			"user":    User.Username,
		})
	}
}
