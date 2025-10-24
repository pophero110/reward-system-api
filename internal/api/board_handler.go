package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"reward-system-api/internal/model"
	"reward-system-api/internal/service"
)

var boardInput struct {
	Name    string `json:"name"`
	OwnerID uint   `json:"owner_id"`
}

func (app *Application) postBoardhandler(w http.ResponseWriter, r *http.Request) {
	if err := json.NewDecoder(r.Body).Decode(&boardInput); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if boardInput.Name == "" || boardInput.OwnerID == 0 {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	board := &model.Board{
		Name:    boardInput.Name,
		OwnerID: boardInput.OwnerID,
	}

	err := app.BoardService.Create(boardInput.Name, board.OwnerID)
	switch {
	case errors.Is(err, service.ErrServerError):
		writeJSONError(w, http.StatusInternalServerError, err.Error())
	case err != nil:
		writeJSONError(w, http.StatusInternalServerError, "unexpected error")
	default:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
	}
}
