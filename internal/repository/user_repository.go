package repository

import (
	"errors"
	"reward-system-api/internal/model"

	"gorm.io/gorm"
)

type UserModel struct {
	DB *gorm.DB
}

// Insert adds a new user to the database
func (m *UserModel) Insert(user *model.User) error {
	if err := m.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

// GetByEmail retrieves a user by email
func (m *UserModel) GetByEmail(email string) (*model.User, error) {
	var u model.User
	err := m.DB.Where("email = ?", email).First(&u).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // no user found
		}
		return nil, err
	}
	return &u, nil
}

// UpdateRole updates a user's role
func (m *UserModel) UpdateRole(userID uint, role string) error {
	return m.DB.Model(&model.User{}).
		Where("id = ?", userID).
		Update("role", role).Error
}
