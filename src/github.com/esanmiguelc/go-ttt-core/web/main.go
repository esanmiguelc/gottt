package main

import (
	"net/http"
	"os"

	"github.com/esanmiguelc/go-ttt-core/web/config"
)

func main() {
	http.ListenAndServe(":"+os.Getenv("PORT"), config.Init())
}
