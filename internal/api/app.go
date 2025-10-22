package api

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"reward-system-api/internal/service"

	_ "github.com/lib/pq" // PostgreSQL driver
)

type Application struct {
	Logger       *slog.Logger
	QuestService *service.QuestService
	UserService  *service.UserService
}

func (app *Application) writeJSON(w http.ResponseWriter, sCode int, data any, headers http.Header) error {
	marshalledJson, err := json.Marshal(data)

	if err != nil {
		return err
	}

	// Valid json requires newline
	marshalledJson = append(marshalledJson, '\n')

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(sCode)
	w.Write(marshalledJson)

	return nil
}
