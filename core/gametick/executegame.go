package gametick

import (
	"errors"
	"strconv"
	"strings"

	"github.com/esanmiguelc/gottt/core/gamestate"
	"github.com/esanmiguelc/gottt/core/rules"
)

func ExecuteGame(playerOneType,
	playerTwoType,
	boardSize,
	movesPlayed string) (gamestate.GameState, bool, string, error) {

	boardSizeAsInt, boardError := strconv.Atoi(boardSize)
	if isValidPlayerType(playerOneType) || isValidPlayerType(playerTwoType) {
		return errorResult()
	}

	if boardError != nil || boardSizeAsInt != rules.THREE_BY_THREE {
		return errorResult()
	}

	if strings.ContainsAny(movesPlayed, "9") {
		return errorResult()
	}
	movesPlayedAsArray := convertMovesPlayed(movesPlayed)
	movesPlayedAsArray = removeDuplicates(movesPlayedAsArray)
	gameState := GameTick(playerOneType, playerTwoType, boardSizeAsInt, movesPlayedAsArray)

	return gameState, rules.IsGameOver(gameState.Board), rules.GetResult(gameState.Board), nil
}

func isValidPlayerType(playerType string) bool {
	return playerType != rules.HUMAN && playerType != rules.COMPUTER
}

func convertMovesPlayed(moves string) []int {
	splitMoves := strings.Split(moves, "")
	var movesPlayed []int
	for _, elem := range splitMoves {
		element, _ := strconv.Atoi(elem)
		movesPlayed = append(movesPlayed, element)
	}

	return movesPlayed
}

func removeDuplicates(moves []int) []int {
	result := []int{}
	seen := map[int]int{}
	for _, value := range moves {
		if _, ok := seen[value]; !ok {
			result = append(result, value)
			seen[value] = value
		}
	}
	return result
}

func errorResult() (gamestate.GameState, bool, string, error) {
	return gamestate.GameState{}, false, "", errors.New("Oops! Something went wrong!")
}
