package api

import (
	"net/http"
)

const version = "0.0.1"

func (app *Application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":  "ok",
		"version": version,
	}

	// No need to add acess control origin headers. On other routes, that may be necessary
	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.Logger.Error(err.Error())
		http.Error(w, "server error", http.StatusInternalServerError)
	}
}
