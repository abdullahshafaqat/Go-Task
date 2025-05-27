package handlers

import (
	"net/http"

	"github.com/abdullahshafaqat/GOTASKS/models"
	"github.com/gin-gonic/gin"
)

func (h *userHandler) LogIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		var entry models.LoginUser
		if err := c.ShouldBindJSON(&entry); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}
		access, refresh, err := h.auth.LoginUser(entry)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"access_token":  access,
			"refresh_token": refresh,
		})
	}
}
