package controllers

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/esanmiguelc/gottt/web/constants"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestGameResponseIsOk(t *testing.T) {
	router := httprouter.New()
	router.GET(constants.GAME_PATH, Game)
	server := httptest.NewServer(router)
	defer server.Close()

	response, _ := http.Get(server.URL + constants.GAME_PATH + "?FirstPlayer=Human&SecondPlayer=Human&BoardSize=3&MovesPlayed=012")
	assert.Equal(t, http.StatusOK, response.StatusCode)
}

func TestGameHasContent(t *testing.T) {
	router := httprouter.New()
	router.GET(constants.GAME_PATH, Game)
	server := httptest.NewServer(router)
	defer server.Close()

	response, _ := http.Get(server.URL + constants.GAME_PATH + "?FirstPlayer=Human&SecondPlayer=Human&BoardSize=3&MovesPlayed=012")

	html, _ := ioutil.ReadAll(response.Body)
	response.Body.Close()
	assert.True(t, strings.Contains(string(html), "Make a Move"))
}

func TestGameRedirectsWhenGameIsOver(t *testing.T) {
	router := httprouter.New()
	router.GET(constants.GAME_PATH, Game)
	router.GET(constants.RESULTS_PATH, Results)
	server := httptest.NewServer(router)
	defer server.Close()

	response, _ := http.Get(server.URL + constants.GAME_PATH + "?FirstPlayer=Human&SecondPlayer=Human&BoardSize=3&MovesPlayed=01438")
	assert.Equal(t, server.URL+constants.RESULTS_PATH+"?Result=X", response.Request.URL.String())
}

func TestIndexResponseIsOk(t *testing.T) {
	router := httprouter.New()
	router.GET(constants.ROOT_PATH, Index)
	server := httptest.NewServer(router)
	defer server.Close()

	response, _ := http.Get(server.URL)
	assert.Equal(t, http.StatusOK, response.StatusCode)
}

func TestIndexHasButtonToStartGame(t *testing.T) {
	router := httprouter.New()
	router.GET(constants.ROOT_PATH, Index)
	server := httptest.NewServer(router)
	defer server.Close()

	response, _ := http.Get(server.URL)
	html, _ := ioutil.ReadAll(response.Body)
	response.Body.Close()
	assert.True(t, strings.Contains(string(html), "Start Game"))
}

func TestResultsPageIsOk(t *testing.T) {
	router := httprouter.New()
	router.GET(constants.RESULTS_PATH, Results)
	server := httptest.NewServer(router)
	defer server.Close()

	response, _ := http.Get(server.URL + constants.RESULTS_PATH + "?result=tie")
	assert.Equal(t, http.StatusOK, response.StatusCode)
}

func TestResultsPageContainsTieResult(t *testing.T) {
	router := httprouter.New()
	router.GET(constants.RESULTS_PATH, Results)
	server := httptest.NewServer(router)
	defer server.Close()

	response, _ := http.Get(server.URL + constants.RESULTS_PATH)
	html, _ := ioutil.ReadAll(response.Body)
	response.Body.Close()
	assert.True(t, strings.Contains(string(html), "tie"))
}

func TestResultsPageContainsWinResult(t *testing.T) {
	router := httprouter.New()
	router.GET(constants.RESULTS_PATH, Results)
	server := httptest.NewServer(router)
	defer server.Close()

	response, _ := http.Get(server.URL + constants.RESULTS_PATH + "?Result=X")
	html, _ := ioutil.ReadAll(response.Body)
	response.Body.Close()
	assert.True(t, strings.Contains(string(html), "X wins"))
}
