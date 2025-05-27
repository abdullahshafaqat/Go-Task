package middlewares

import (
	"errors"
	"regexp"

	"github.com/abdullahshafaqat/GOTASKS/models"
	"golang.org/x/crypto/bcrypt"
)

func (a *authMiddleware) RegisterUser(user models.NewUser) error {
	if !isValidGmail(user.Email) {
		return errors.New("invalid email")
	}
	exists, err := a.repo.GetUser(user.Username, user.Email)
	if err == nil && exists {
		return errors.New("user already exists")
	}
	return a.repo.CreateUser(user)
}
func isValidGmail(email string) bool {
	valid := regexp.MustCompile(`^[^@]+@gmail\.com$`)
	return valid.MatchString(email)
}

func (a *authMiddleware) LoginUser(entry models.LoginUser) (string, string, error) {
	info, err := a.repo.GetUserByEmail(entry.Email)
	if err != nil {
		return "", "", errors.New("invalid email or password")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(info.Password), []byte(entry.Password)); err != nil {
		return "", "", errors.New("invalid password")
	}
	access, refresh, err := a.GenerateTokens(info.ID)
	if err != nil {
		return "", "", errors.New("failed to generate tokens")
	}
	return access, refresh, err
}
