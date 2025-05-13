package main

import (
	"github.com/abdullahshafaqat/GOTASKS/database"
	"github.com/abdullahshafaqat/GOTASKS/handlers"
	"github.com/gin-gonic/gin"
)

func main() {

	db := database.Database()
	defer db.Close()
	router := gin.Default()
	router.POST("/analyzer", handlers.AnalyzeText(db))


	router.POST("/signup", handlers.SignUp(db))

	router.POST("/login",handlers.LogIn(db))
	
	router.Run(":8080")
}
