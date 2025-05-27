package repo

import (
	"github.com/abdullahshafaqat/GOTASKS/models"
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	CreateUser(user models.NewUser) error
	GetUser(username, email string) (bool, error)
	GetUserByEmail(email string) (models.NewUser, error)
}

type UserRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepository {
	return &UserRepo{
		db: db,
	}
}