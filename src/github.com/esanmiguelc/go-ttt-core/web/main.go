package main

import (
	"net/http"

	"github.com/esanmiguelc/go-ttt-core/web/config"
)

func main() {
	http.ListenAndServe(":8080", config.Init())
}
