package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *userHandler) RefreshToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		refreshToken := c.PostForm("refresh_token")
		if refreshToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token required"})
			return
		}

		email, err := h.auth.VerifyRefreshToken(refreshToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
			return
		}
		newAccessToken, _, err := h.auth.GenerateTokens(email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate new token"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"access_token": newAccessToken,
		})
	}
}
