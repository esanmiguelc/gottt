package controllers

import (
	"net/http"
	"path"
	"runtime"
	"strings"
	"text/template"

	"github.com/esanmiguelc/gottt/web/constants"
	"github.com/esanmiguelc/gottt/web/interactors"
	"github.com/esanmiguelc/gottt/web/viewmodels"
	"github.com/julienschmidt/httprouter"
)

func Error(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.WriteHeader(500)
	render("error", writer, nil)
}

func Game(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gameState, isGameOver, result, err := interactors.ExecuteGameInteractor(
		request.FormValue("FirstPlayer"),
		request.FormValue("SecondPlayer"),
		request.FormValue("BoardSize"),
		request.FormValue("MovesPlayed"),
	)

	if err != nil {
		http.Redirect(writer, request, constants.ERROR_PATH, http.StatusFound)
	}

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
	if strings.ToLower(result) == strings.ToLower("Tie") {
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
