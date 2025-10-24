package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"reward-system-api/internal/service"
)

type registerInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (app *Application) registerHandler(w http.ResponseWriter, r *http.Request) {
	var input registerInput

	// Parse JSON body
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Validate input
	if input.Email == "" || input.Password == "" {
		http.Error(w, "Missing email or password", http.StatusBadRequest)
		return
	}

	// Create user
	err := app.UserService.Create(input.Email, input.Password)
	switch {
	case errors.Is(err, service.ErrUserExists):
		writeJSONError(w, http.StatusBadRequest, err.Error())
	case errors.Is(err, service.ErrServerError):
		writeJSONError(w, http.StatusInternalServerError, err.Error())
	case err != nil:
		writeJSONError(w, http.StatusInternalServerError, "unexpected error")
	default:
		writeJSON(w, http.StatusCreated, nil)
	}
}
