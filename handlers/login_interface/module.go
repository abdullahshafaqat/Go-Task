package login

import (
	"github.com/abdullahshafaqat/GOTASKS/models"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type UserLoginRepo interface {
	GetUserByEmail(db *sqlx.DB, email string) (models.NewUser, error)
}

type UserLoginService interface {
	LoginUser(db *sqlx.DB, entry models.LoginUser) (string, string, error)
}

type UserLoginHandler interface {
	LogIn(db *sqlx.DB) func(c *gin.Context)
}
