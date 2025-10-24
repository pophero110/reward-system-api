package service

import (
	"log/slog"
	"reward-system-api/internal/model"
	"reward-system-api/internal/repository"
)

type BoardService struct {
	Logger *slog.Logger
	Boards *repository.BoardModel
}

func (service *BoardService) Create(name string, ownerID uint) error {
	board := model.Board{
		Name:    name,
		OwnerID: ownerID,
	}
	if err := service.Boards.Insert(&board); err != nil {
		return ErrServerError
	}
	return nil
}
