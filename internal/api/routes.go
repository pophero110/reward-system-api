package api

import (
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

var tmpl = template.Must(template.ParseFiles("templates/index.html"))

func (app *Application) indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html", "templates/quests.html"))
	tmpl.ExecuteTemplate(w, "index.html", nil)
}

func (app *Application) questsListHandler(w http.ResponseWriter, r *http.Request) {
	quests, err := app.QuestService.GetAll()
	if err != nil {
		http.Error(w, "failed to fetch quests", http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/quest_item.html"))
	for _, q := range quests {
		tmpl.ExecuteTemplate(w, "quest_item.html", q)
	}
}

func (app *Application) Routes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/", app.indexHandler)
	router.HandlerFunc(http.MethodGet, "/quests", app.questsListHandler)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	// Quests
	router.HandlerFunc(http.MethodGet, "/v1/quests", app.getQuestHandler)
	router.HandlerFunc(http.MethodPost, "/v1/quests", app.postQuestHandler)
	// Uesrs
	router.HandlerFunc(http.MethodPost, "/v1/users", app.registerHandler)
	// Boards
	router.HandlerFunc(http.MethodPost, "/v1/boards", app.postBoardhandler)

	return router
}
