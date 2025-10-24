package service

import (
	"log/slog"
	"reward-system-api/internal/model"
	"reward-system-api/internal/repository"
)

type QuestService struct {
	Logger *slog.Logger
	Quests *repository.QuestModel
}

func (service *QuestService) GetAll() ([]*model.Quest, error) {
	quests, err := service.Quests.GetAll()
	if err != nil {
		service.Logger.Error("failed to fetch quests", "error", err)
		return nil, err
	}

	return quests, nil
}

func (service *QuestService) Create(quest *model.Quest) error {
	if err := service.Quests.Insert(quest); err != nil {
		return err
	}
	return nil
}
