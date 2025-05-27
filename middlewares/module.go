package middlewares

import (
	repo "github.com/abdullahshafaqat/GOTASKS/api/repository"
	"github.com/abdullahshafaqat/GOTASKS/models"
	"github.com/jmoiron/sqlx"
)

type AuthMiddleware interface {
	RegisterUser(user models.NewUser) error
	LoginUser(entry models.LoginUser) (string, string, error)
	GenerateTokens(ID string) (string, string, error)
	VerifyToken(tokenString string) error
	VerifyRefreshToken(tokenString string) (string, error)
}

type authMiddleware struct {
	repo repo.UserRepository
}

func NewAuthMiddleware(db *sqlx.DB) AuthMiddleware {
	return &authMiddleware{
		repo: repo.NewUserRepo(db),
	}
}
