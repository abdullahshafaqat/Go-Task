package repositories

import (
	"github.com/abdullahshafaqat/GOTASKS/models"
	"golang.org/x/crypto/bcrypt"
)

func HashedPassword(User models.NewUser) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(User.Password), bcrypt.DefaultCost)
	return string(hashed), err
}
