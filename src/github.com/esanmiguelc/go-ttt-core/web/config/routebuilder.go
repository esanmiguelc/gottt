package config

import (
	"github.com/esanmiguelc/go-ttt-core/Godeps/_workspace/src/github.com/julienschmidt/httprouter"
	"github.com/esanmiguelc/go-ttt-core/web/constants"
	"github.com/esanmiguelc/go-ttt-core/web/controllers"
)

func Init() *httprouter.Router {
	router := httprouter.New()
	router.GET(constants.ROOT_PATH, controllers.Index)
	router.GET(constants.GAME_PATH, controllers.Game)
	router.GET(constants.RESULTS_PATH, controllers.Results)
	return router
}
