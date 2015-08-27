package controllers

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/esanmiguelc/go-ttt-core/web/config"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestGameResponseIsOk(t *testing.T) {
	router := httprouter.New()
	router.GET(config.GAME_PATH, Game)
	server := httptest.NewServer(router)
	defer server.Close()

	response, _ := http.Get(server.URL + config.GAME_PATH)
	assert.Equal(t, http.StatusOK, response.StatusCode)
}

func TestGameHasContent(t *testing.T) {
	router := httprouter.New()
	router.GET(config.GAME_PATH, Game)
	server := httptest.NewServer(router)
	defer server.Close()

	response, _ := http.Get(server.URL + config.GAME_PATH)
	html, _ := ioutil.ReadAll(response.Body)
	response.Body.Close()
	assert.True(t, strings.Contains(string(html), "Make a Move"))
}

func TestGameRedirectsWhenGameIsOver(t *testing.T) {
	router := httprouter.New()
	router.GET(config.GAME_PATH, Game)
	router.GET(config.RESULTS_PATH, Results)
	server := httptest.NewServer(router)
	defer server.Close()

	response, _ := http.Get(server.URL + config.GAME_PATH + "?FirstPlayer=1&SecondPlayer=1&GameState=012345678")
	assert.Equal(t, server.URL+config.RESULTS_PATH, response.Request.URL.String())
}

func TestIndexResponseIsOk(t *testing.T) {
	router := httprouter.New()
	router.GET(config.ROOT_PATH, Index)
	server := httptest.NewServer(router)
	defer server.Close()

	response, _ := http.Get(server.URL)
	assert.Equal(t, http.StatusOK, response.StatusCode)
}

func TestIndexHasButtonToStartGame(t *testing.T) {
	router := httprouter.New()
	router.GET(config.ROOT_PATH, Index)
	server := httptest.NewServer(router)
	defer server.Close()

	response, _ := http.Get(server.URL)
	html, _ := ioutil.ReadAll(response.Body)
	response.Body.Close()
	assert.True(t, strings.Contains(string(html), "Start Game"))
}

func TestResultsPageIsOk(t *testing.T) {
	router := httprouter.New()
	router.GET(config.RESULTS_PATH, Results)
	server := httptest.NewServer(router)
	defer server.Close()

	response, _ := http.Get(server.URL + config.RESULTS_PATH + "?result=tie")
	assert.Equal(t, http.StatusOK, response.StatusCode)
}

func TestResultsPageContainsTieResult(t *testing.T) {
	router := httprouter.New()
	router.GET(config.RESULTS_PATH, Results)
	server := httptest.NewServer(router)
	defer server.Close()

	response, _ := http.Get(server.URL + config.RESULTS_PATH)
	html, _ := ioutil.ReadAll(response.Body)
	response.Body.Close()
	assert.True(t, strings.Contains(string(html), "tie"))
}

func TestResultsPageContainsWinResult(t *testing.T) {
	router := httprouter.New()
	router.GET(config.RESULTS_PATH, Results)
	server := httptest.NewServer(router)
	defer server.Close()

	response, _ := http.Get(server.URL + config.RESULTS_PATH + "?Result=X")
	html, _ := ioutil.ReadAll(response.Body)
	response.Body.Close()
	assert.True(t, strings.Contains(string(html), "X wins"))
}
