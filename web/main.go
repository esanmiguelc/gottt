package main

import (
	"net/http"
	"os"

	"github.com/esanmiguelc/gottt/web/config"
)

func main() {
	http.ListenAndServe(":"+os.Getenv("PORT"), config.Init())
}
