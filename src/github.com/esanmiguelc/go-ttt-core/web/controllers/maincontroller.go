package controllers

import (
	"net/http"
	"strings"
	"text/template"

	"github.com/esanmiguelc/go-ttt-core/core"
	"github.com/esanmiguelc/go-ttt-core/web/config"
	"github.com/esanmiguelc/go-ttt-core/web/viewmodels"
	"github.com/julienschmidt/httprouter"
)

func Game(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// BuildGame(request.FormValue("FirstPlayer"), request.FormValue("SecondPlayer"), request.FormValue("GameState"))
	requestFormState := request.FormValue("GameState")
	splitState := strings.Split(requestFormState, "")
	for index := range splitState {
		splitState[index] = strings.Replace(splitState[index], " ", "", -1)
	}
	board := core.Board{Slots: splitState}
	gameviewmodel := viewmodels.GameViewModel{State: splitState, PlayerOne: 1, PlayerTwo: 2}
	if core.IsGameOver(board, "X", "O") {
		http.Redirect(writer, request, config.RESULTS_PATH, http.StatusFound)
	} else {
		render("game", writer, gameviewmodel)
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
