package repo

import (
	signup "github.com/abdullahshafaqat/GOTASKS/handlers/signup_interface"
	"github.com/abdullahshafaqat/GOTASKS/models"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type userRepository struct{}
func NewUser() signup.UserSignupRepo {
	return &userRepository{}

}

// For signup
func (r *userRepository) HashedPassword(User models.NewUser) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(User.Password), bcrypt.DefaultCost)
	return string(hashed), err
}
func (r *userRepository) GetUser(db *sqlx.DB, username, email string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE username = $1 OR email = $2)`
	err := db.QueryRow(query, username, email).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (r *userRepository) CreateUser(db *sqlx.DB, User models.NewUser) error {
	hashed, err := r.HashedPassword(User)
	if err != nil {
		return err
	}
	query := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id`
	_ = db.QueryRow(query, User.Username, User.Email, hashed)
	return err
}

// For login

func GetUserByEmail(db *sqlx.DB, email string) (models.NewUser, error) {
	var info models.NewUser
	query := `SELECT id, username, email, password FROM users WHERE email = $1`
	err := db.Get(&info, query, email)
	return info, err

}
var _ signup.UserSignupRepo = &userRepository{}
