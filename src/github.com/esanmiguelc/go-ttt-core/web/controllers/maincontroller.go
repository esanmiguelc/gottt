package controllers

import (
	"net/http"
	"text/template"

	"github.com/esanmiguelc/go-ttt-core/core"
	"github.com/esanmiguelc/go-ttt-core/web/constants"
	"github.com/esanmiguelc/go-ttt-core/web/viewmodels"
	"github.com/julienschmidt/httprouter"
)

func Game(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	playerOneType := request.FormValue("PlayerOne")
	playerTwoType := request.FormValue("PlayerTwo")
	boardSize := request.FormValue("BoardSize")
	movesPlayed := request.FormValue("MovesPlayed")
	gameState := core.GameTick(playerOneType, PlayerTwoType, boardSize, movesPlayed)
	if core.IsGameOver(gameState) {
		http.Redirect(writer, request, constants.RESULTS_PATH, http.StatusFound)
	} else {
		render("game", writer, gameState)
	}
}

func Index(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	render("index", writer, nil)
}

func Results(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	viewmodel := viewmodels.ResultsViewModel{Result: request.FormValue("Result")}
	render("results", writer, viewmodel)
}

func render(viewName string, writer http.ResponseWriter, model interface{}) {
	templates := template.Must(template.ParseGlob("../views/*"))
	templates.ExecuteTemplate(writer, viewName, model)
}
