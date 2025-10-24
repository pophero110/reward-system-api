package model

import (
	"time"

	"gorm.io/gorm"
)

type QuestStatus string

const (
	StatusOpen       QuestStatus = "open"
	StatusInProgress QuestStatus = "in_progress"
	StatusCompleted  QuestStatus = "completed"
	StatusCancelled  QuestStatus = "cancelled"
)

// Quest represents the "quest" table in the database
type Quest struct {
	gorm.Model
	BoardID     uint
	AssigneeID  *uint
	CreatorID   uint
	Title       string
	Description string
	Status      QuestStatus
	Reward      string
	DueDateTime *time.Time

	// Association
	Creator  User
	Assignee User
	Board    Board
}
