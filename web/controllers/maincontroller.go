package controllers

import (
	"net/http"
	"path"
	"runtime"
	"strings"
	"text/template"

	"github.com/esanmiguelc/gottt/core/gametick"
	"github.com/esanmiguelc/gottt/core/rules"
	"github.com/esanmiguelc/gottt/web/constants"
	"github.com/esanmiguelc/gottt/web/validators"
	"github.com/esanmiguelc/gottt/web/viewmodels"
	"github.com/julienschmidt/httprouter"
)

func Error(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.WriteHeader(http.StatusInternalServerError)
	render("error", writer, nil)
}

func Game(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	firstPlayerType, secondPlayerType, boardSize, movesPlayed, err := validators.ValidateParams(
		request.FormValue("FirstPlayer"),
		request.FormValue("SecondPlayer"),
		request.FormValue("BoardSize"),
		request.FormValue("MovesPlayed"),
	)

	if err != nil {
		http.Redirect(writer, request, constants.ERROR_PATH, http.StatusFound)
	}

	gameState, isGameOver, result := gametick.ExecuteGame(firstPlayerType, secondPlayerType, boardSize, movesPlayed)

	if isGameOver {
		http.Redirect(writer, request, constants.RESULTS_PATH+"?Result="+result, http.StatusFound)
	} else {
		render("game", writer, gameState)
	}
}

func Index(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	render("index", writer, nil)
}

func Results(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	result := request.FormValue("Result")
	var viewmodel viewmodels.ResultsViewModel
	if strings.ToLower(result) == strings.ToLower(rules.TIE) {
		viewmodel = viewmodels.ResultsViewModel{Result: result, Tie: true}
	} else {
		viewmodel = viewmodels.ResultsViewModel{Result: result, Tie: false}
	}
	render("results", writer, viewmodel)
}

func render(viewName string, writer http.ResponseWriter, model interface{}) {
	_, filename, _, _ := runtime.Caller(1)
	f := path.Join(path.Dir(filename), "../views/*")
	templates := template.Must(template.ParseGlob(f))
	templates.ExecuteTemplate(writer, viewName, model)
}
