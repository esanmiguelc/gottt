package config

import (
	"github.com/esanmiguelc/gottt/web/constants"
	"github.com/esanmiguelc/gottt/web/controllers"
	"github.com/julienschmidt/httprouter"
)

func Init() *httprouter.Router {
	router := httprouter.New()
	router.GET(constants.ROOT_PATH, controllers.Index)
	router.GET(constants.GAME_PATH, controllers.Game)
	router.GET(constants.RESULTS_PATH, controllers.Results)
	return router
}
