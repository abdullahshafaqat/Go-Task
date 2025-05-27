package handlers

import (
	"github.com/abdullahshafaqat/GOTASKS/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type UserHandler interface {
	SignUp() gin.HandlerFunc
	LogIn() gin.HandlerFunc
	RefreshToken() gin.HandlerFunc
	Authorization() gin.HandlerFunc
}

type userHandler struct {
	db   *sqlx.DB
	auth middlewares.AuthMiddleware
}

func NewUserHandler(db *sqlx.DB) UserHandler {
	return &userHandler{
		db:   db,
		auth: middlewares.NewAuthMiddleware(db), // create auth middleware
	}
}
