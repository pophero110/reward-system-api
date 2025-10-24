package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"reward-system-api/internal/model"
	"time"
)

// Input DTO
type QuestInput struct {
	BoardID     uint              `json:"board_id"`
	CreatorID   uint              `json:"creator_id"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	DueDateTime *time.Time        `json:"due_date_time"`
	Reward      string            `json:"reward"`
	Status      model.QuestStatus `json:"status"`
}

// Output DTO
type QuestDto struct {
	ID          uint              `json:"id"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	DueDateTime *time.Time        `json:"due_date_time"`
	Reward      string            `json:"reward"`
	Status      model.QuestStatus `json:"status"`
}

// ----------------------
// Helper Functions
// ----------------------

// mapQuestToDTO converts model.Quest to QuestDto
func mapQuestToDTO(q *model.Quest) QuestDto {
	return QuestDto{
		ID:          q.ID,
		Title:       q.Title,
		Description: q.Description,
		DueDateTime: q.DueDateTime,
		Reward:      q.Reward,
		Status:      q.Status,
	}
}

// mapQuestsToDTOs converts a slice of model.Quest to []QuestDto
func mapQuestsToDTOs(quests []*model.Quest) []QuestDto {
	dtos := make([]QuestDto, len(quests))
	for i, q := range quests {
		dtos[i] = mapQuestToDTO(q)
	}
	return dtos
}

// validateQuestInput ensures required fields are present
func validateQuestInput(input *QuestInput) error {
	if input.Title == "" || input.BoardID == 0 {
		return errors.New("Invalid input")
	}
	return nil
}

// ----------------------
// Handlers
// ----------------------

// GET /quests
func (app *Application) getQuestHandler(w http.ResponseWriter, r *http.Request) {
	quests, err := app.QuestService.GetAll()
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, "failed to fetch quests")
		return
	}

	writeJSON(w, http.StatusOK, mapQuestsToDTOs(quests))
}

// POST /quests
func (app *Application) postQuestHandler(w http.ResponseWriter, r *http.Request) {
	var input QuestInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeJSONError(w, http.StatusBadRequest, "invalid JSON")
		return
	}

	if err := validateQuestInput(&input); err != nil {
		writeJSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	quest := &model.Quest{
		BoardID:     input.BoardID,
		CreatorID:   input.CreatorID,
		Title:       input.Title,
		Description: input.Description,
		DueDateTime: input.DueDateTime,
		Reward:      input.Reward,
		Status:      input.Status,
	}

	if err := app.QuestService.Create(quest); err != nil {
		app.Logger.Error("failed to insert quest", "error", err)
		writeJSONError(w, http.StatusInternalServerError, "database error")
		return
	}

	writeJSON(w, http.StatusCreated, mapQuestToDTO(quest))
}
