package main

import (
	database "github.com/abdullahshafaqat/GOTASKS/Database"
	"github.com/abdullahshafaqat/GOTASKS/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	db := database.Database()
	defer db.Close()
	router := gin.Default()
	routes.DefineRoutes(router, db)
	router.Run(":8080")
}
