package model

import "gorm.io/gorm"

type Board struct {
	gorm.Model
	Name    string
	OwnerID uint
	// Association
	Owner   User
	Members []User
	Quests  []Quest
}
