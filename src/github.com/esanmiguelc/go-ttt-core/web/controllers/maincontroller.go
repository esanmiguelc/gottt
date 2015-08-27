package controllers

import (
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

func Index(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	render("index.html", writer, nil)
}

func Game(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	render("game.html", writer, nil)
}

func render(viewName string, writer http.ResponseWriter, model interface{}) {
	t, _ := template.ParseFiles("../views/" + viewName)
	t.Execute(writer, model)
}
