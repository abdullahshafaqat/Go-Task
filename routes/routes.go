package routes

import (
	"github.com/abdullahshafaqat/GOTASKS/handlers"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func DefineRoutes(router *gin.Engine, db *sqlx.DB) {
	router.POST("/analyzer", handlers.Authorization(db), handlers.UploadText(db))
	router.POST("/signup", (handlers.NewUserHandler().SignUp(db)))
	router.POST("/login", handlers.LogIn(db))
	router.POST("/auth", handlers.Authorization(db))
	router.POST("/refresh", handlers.RefreshToken)
}
