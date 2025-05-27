package routes

import (
	"github.com/abdullahshafaqat/GOTASKS/handlers"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func DefineRoutes(router *gin.Engine, db *sqlx.DB) {
	handler := handlers.NewUserHandler(db)
	router.POST("/analyzer", handler.Authorization(), handlers.UploadText(db))
	router.POST("/signup", handler.SignUp())
	router.POST("/login", handler.LogIn())
	router.POST("/auth", handler.Authorization())
	router.POST("/refresh", handler.RefreshToken())
}
