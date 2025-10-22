package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *Application) Routes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	// Quests
	router.HandlerFunc(http.MethodGet, "/v1/quests", app.getAllQuests)
	router.HandlerFunc(http.MethodPost, "/v1/quests", app.postQuestHandler)
	// Uesrs
	router.HandlerFunc(http.MethodPost, "/v1/users", app.registerHandler)

	return router
}
