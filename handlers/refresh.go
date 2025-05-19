package handlers

import (
	"net/http"

	"github.com/abdullahshafaqat/GOTASKS/services"
	"github.com/gin-gonic/gin"
)

func RefreshToken(c *gin.Context) {
	refreshToken := c.PostForm("refresh_token")
	if refreshToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token required"})
		return
	}

	email, err := services.VerifyRefreshToken(refreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}
	newAccessToken,_,err := services.GenerateTokens(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate new token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"access_token": newAccessToken,
	})
}
