package services

import (
	"errors"
	"regexp"

	repo "github.com/abdullahshafaqat/GOTASKS/api/repository"
	"github.com/abdullahshafaqat/GOTASKS/models"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

func isValidGmail(email string) bool {
	valid := regexp.MustCompile(`^[^@]+@gmail\.com$`)
	return valid.MatchString(email)
}

func RegisterUser(db *sqlx.DB, User models.NewUser) error {
	if !isValidGmail(User.Email) {
		return errors.New("invalid email")
	}
	exists, err := repo.GetUser(db, User.Username, User.Email)
	if err == nil && exists {
		return errors.New("user already exists")
	}
	return repo.CreateUser(db, User)
}

func LoginUser(db *sqlx.DB, entry models.LoginUser) (string, string, error) {
	info, err := repo.GetUserByEmail(db, entry.Email)
	if err != nil {
		return "", "", errors.New("invalid email or password")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(info.Password), []byte(entry.Password)); err != nil {
		return "", "", errors.New("invalid password")
	}
	access, refresh, err:= GenerateTokens(info.ID)
	if err != nil {
		return "", "", errors.New("failed to generate tokens")
	}
	return access, refresh, err
}
