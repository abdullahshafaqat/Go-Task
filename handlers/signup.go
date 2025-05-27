package handlers

import (
	"net/http"

	"github.com/abdullahshafaqat/GOTASKS/middlewares"
	"github.com/abdullahshafaqat/GOTASKS/models"

	"github.com/gin-gonic/gin"
)

func (h *userHandler) SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var User models.NewUser

		if err := c.ShouldBindJSON(&User); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	auth := middlewares.NewAuthMiddleware(h.db)
	err := auth.RegisterUser(User)
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
