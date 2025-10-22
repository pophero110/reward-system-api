package api

import (
	"encoding/json"
	"net/http"
	"reward-system-api/internal/model"
	"time"
)

var questInput struct {
	BoardID     uint       `json:"board_id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	DueDateTime *time.Time `json:"due_date_time"`
	Comments    string     `json:"comments"`
	Reward      string     `json:"reward"`
	Status      string     `json:"status"`
}

// GET /quests
func (app *Application) getAllQuests(w http.ResponseWriter, r *http.Request) {
	quests, err := app.QuestService.GetAll()
	if err != nil {
		app.Logger.Error("failed to fetch quests", "error", err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(quests); err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
	}
}

// POST /quests
func (app *Application) postQuestHandler(w http.ResponseWriter, r *http.Request) {
	if err := json.NewDecoder(r.Body).Decode(&questInput); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if questInput.Title == "" || questInput.BoardID == 0 {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	quest := &model.Quest{
		BoardID:     questInput.BoardID,
		Title:       questInput.Title,
		Description: questInput.Description,
		Reward:      questInput.Reward,
		Status:      questInput.Status,
	}

	if err := app.QuestService.Create(*quest); err != nil {
		app.Logger.Error("failed to insert quest", "error", err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(quest)
}
