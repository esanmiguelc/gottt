package controllers

import (
	"net/http"
	"text/template"

	"github.com/esanmiguelc/go-ttt-core/web/config"
	"github.com/julienschmidt/httprouter"
)

type ResultsViewModel struct {
	Result string
}

func Game(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	if request.FormValue("GameState") == "012345678" {
		http.Redirect(writer, request, config.RESULTS_PATH, http.StatusFound)
	} else {
		render("game", writer, nil)
	}
}

func Index(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	render("index", writer, nil)
}

func Results(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	viewmodel := ResultsViewModel{Result: request.FormValue("Result")}
	render("results", writer, viewmodel)
}

func render(viewName string, writer http.ResponseWriter, model interface{}) {
	templates := template.Must(template.ParseGlob("../views/*"))
	templates.ExecuteTemplate(writer, viewName, model)
}
