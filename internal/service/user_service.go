package service

import (
	"log/slog"
	"reward-system-api/internal/model"
	"reward-system-api/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Logger *slog.Logger
	Users  *repository.UserModel
}

func (service *UserService) Create(email string, password string) error {
	existing, err := service.Users.GetByEmail(email)
	if err != nil {
		return ErrServerError
	}
	if existing != nil {
		return ErrUserExists
	}

	hashed, err := hashPassword(password)
	if err != nil {
		return ErrServerError
	}

	user := &model.User{
		Email:          email,
		HashedPassword: string(hashed),
		Role:           "user", // default role
	}

	if err := service.Users.Insert(user); err != nil {
		return ErrServerError
	}
	return nil
}

// HashPassword hashes a plain password
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPassword verifies a plain password against a hash
func checkPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
