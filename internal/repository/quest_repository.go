package repository

import (
	"gorm.io/gorm"
	"reward-system-api/internal/model"
)

type QuestModel struct {
	DB *gorm.DB
}

// GetAll fetches all quests from the database (ordered by due date)
func (m *QuestModel) GetAll() ([]*model.Quest, error) {
	var quests []*model.Quest

	if err := m.DB.
		Order("due_date_time ASC").
		Find(&quests).Error; err != nil {
		return nil, err
	}

	return quests, nil
}

// Insert adds a new quest to the database and updates its ID and timestamps
func (m *QuestModel) Insert(q *model.Quest) error {
	if err := m.DB.Create(q).Error; err != nil {
		return err
	}
	return nil
}
