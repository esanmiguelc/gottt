package main

import (
	"net/http"

	"github.com/esanmiguelc/go-ttt-core/web/controllers"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/", controllers.Index)
	router.GET("/game", controllers.Game)
	router.GET("/results", controllers.Results)
	http.ListenAndServe(":8080", router)
}
