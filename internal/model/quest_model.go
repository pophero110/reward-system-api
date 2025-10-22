package model

import (
	"time"

	"gorm.io/gorm"
)

// Quest represents the "quest" table in the database
type Quest struct {
	gorm.Model
	BoardID     uint
	AssigneeID  *uint
	CreatorID   uint
	Title       string
	Description string
	Status      string
	Reward      string
	DueDateTime *time.Time

	// Association
	Creator  User
	Assignee User
	Board    Board
}
