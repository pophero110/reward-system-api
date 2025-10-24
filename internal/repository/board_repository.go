package repository

import (
	"gorm.io/gorm"
	"reward-system-api/internal/model"
)

type BoardModel struct {
	DB *gorm.DB
}

func (m *BoardModel) Insert(q *model.Board) error {
	if err := m.DB.Create(q).Error; err != nil {
		return err
	}
	return nil
}
