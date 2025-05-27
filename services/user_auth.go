package services

import (
	"errors"
	"regexp"

	repo "github.com/abdullahshafaqat/GOTASKS/api/repository"
	signup "github.com/abdullahshafaqat/GOTASKS/handlers/signup_interface"
	"github.com/abdullahshafaqat/GOTASKS/models"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repo signup.UserSignupRepo
}

func NewUserService() signup.UserSignupService {
	return &userService{
		repo: repo.NewUser(),
	}
}

func (s *userService) RegisterUser(db *sqlx.DB, user models.NewUser) error {
	if !s.isValidGmail(user.Email) {
		return errors.New("invalid email")
	}
	exists, err := s.repo.GetUser(db, user.Username, user.Email)
	if err == nil && exists {
		return errors.New("user already exists")
	}
	return s.repo.CreateUser(db, user)
}
func (s *userService) isValidGmail(email string) bool {
	valid := regexp.MustCompile(`^[^@]+@gmail\.com$`)
	return valid.MatchString(email)
}

func LoginUser(db *sqlx.DB, entry models.LoginUser) (string, string, error) {
	info, err := repo.GetUserByEmail(db, entry.Email)
	if err != nil {
		return "", "", errors.New("invalid email or password")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(info.Password), []byte(entry.Password)); err != nil {
		return "", "", errors.New("invalid password")
	}
	access, refresh, err := GenerateTokens(info.ID)
	if err != nil {
		return "", "", errors.New("failed to generate tokens")
	}
	return access, refresh, err
}

var _ signup.UserSignupService = &userService{}
