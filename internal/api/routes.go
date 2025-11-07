package api

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
)

// HTML Response Experiment
func (app *Application) indexHandler(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseFiles("templates/layout.html", "templates/quests.html"))
	tmpl.ExecuteTemplate(w, "layout.html", "This is layout html")
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

	router.HandlerFunc(http.MethodGet, "/healthcheck", app.healthcheckHandler)
	// Quests
	router.HandlerFunc(http.MethodPost, "/quests", app.postQuestHandler)
	router.HandlerFunc(http.MethodGet, "/quests", app.getQuestsHandler)
	router.GET("/quests/:id", app.getQuestByIdHandler)
	// Uesrs
	router.HandlerFunc(http.MethodPost, "/users", app.registerHandler)
	// Boards
	router.HandlerFunc(http.MethodPost, "/boards", app.postBoardhandler)

	return router
}
