package controllers

import (
	"net/http"
	"path"
	"runtime"
	"strconv"
	"strings"
	"text/template"

	"github.com/esanmiguelc/gottt/core"
	"github.com/esanmiguelc/gottt/web/constants"
	"github.com/esanmiguelc/gottt/web/viewmodels"
	"github.com/julienschmidt/httprouter"
)

func Game(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	playerOneType := request.FormValue("FirstPlayer")
	playerTwoType := request.FormValue("SecondPlayer")
	boardSize, _ := strconv.Atoi(request.FormValue("BoardSize"))
	someVar := strings.Split(request.FormValue("MovesPlayed"), "")
	var movesPlayed []int
	for _, elem := range someVar {
		element, _ := strconv.Atoi(elem)
		movesPlayed = append(movesPlayed, element)
	}
	gameState := core.GameTick(playerOneType, playerTwoType, boardSize, movesPlayed)
	if core.IsGameOver(gameState.Board) {
		if core.IsWinner(gameState.Board, core.FIRST_PLAYER) {
			http.Redirect(writer, request, constants.RESULTS_PATH+"?Result="+core.FIRST_PLAYER, http.StatusFound)
		} else if core.IsWinner(gameState.Board, core.SECOND_PLAYER) {
			http.Redirect(writer, request, constants.RESULTS_PATH+"Result="+core.SECOND_PLAYER, http.StatusFound)
		} else {
			http.Redirect(writer, request, constants.RESULTS_PATH, http.StatusFound)
		}
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
	_, filename, _, _ := runtime.Caller(1)
	f := path.Join(path.Dir(filename), "../views/*")
	templates := template.Must(template.ParseGlob(f))
	templates.ExecuteTemplate(writer, viewName, model)
}
