package interactors

import (
	"errors"
	"strconv"
	"strings"

	"github.com/esanmiguelc/gottt/core"
)

func ExecuteGameInteractor(playerOneType,
	playerTwoType,
	boardSize,
	movesPlayed string) (core.GameState, bool, string, error) {

	boardSizeAsInt, boardError := strconv.Atoi(boardSize)
	if isValidPlayerType(playerOneType) || isValidPlayerType(playerTwoType) {
		return errorResult()
	}

	if boardError != nil || boardSizeAsInt != core.THREE_BY_THREE {
		return errorResult()
	}

	if strings.ContainsAny(movesPlayed, "9") {
		return errorResult()
	}

	movesPlayedAsArray := convertMovesPlayed(movesPlayed)
	gameState := core.GameTick(playerOneType, playerTwoType, boardSizeAsInt, movesPlayedAsArray)

	if core.IsGameOver(gameState.Board) {
		return gameState, true, core.GetResult(gameState.Board), nil
	} else {
		return gameState, false, "", nil
	}
}

func isValidPlayerType(playerType string) bool {
	return playerType != core.HUMAN && playerType != core.COMPUTER
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

func errorResult() (core.GameState, bool, string, error) {
	return core.GameState{}, false, "", errors.New("Oops! Something went wrong!")
}
