package signup

import (
	"github.com/abdullahshafaqat/GOTASKS/models"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type UserSignupRepo interface {
	HashedPassword(User models.NewUser) (string, error)
	GetUser(db *sqlx.DB, username, email string) (bool, error)
	CreateUser(db *sqlx.DB, User models.NewUser) error
}

type UserSignupService interface {
	RegisterUser(db *sqlx.DB, User models.NewUser) error
}

type UserSignupHandler interface {
	SignUp(db *sqlx.DB) gin.HandlerFunc
}
