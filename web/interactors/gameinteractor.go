package interactors

import (
	"errors"
	"strconv"
	"strings"

	"github.com/esanmiguelc/gottt/core"
)

var emptyState, gameErrors = core.GameState{}, errors.New("Oops! Something went wrong parsing the board size")

func ExecuteGameInteractor(playerOneType,
	playerTwoType,
	boardSize,
	movesPlayed string) (core.GameState, error) {
	boardSizeAsInt, boardError := strconv.Atoi(boardSize)
	if isValidPlayerType(playerOneType) || isValidPlayerType(playerTwoType) {
		return emptyState, gameErrors
	}
	if boardError != nil || boardSizeAsInt != core.THREE_BY_THREE {
		return emptyState, gameErrors
	}
	movesPlayedAsArray := convertMovesPlayed(movesPlayed)
	gameState := core.GameTick(playerOneType, playerTwoType, boardSizeAsInt, movesPlayedAsArray)
	return gameState, nil
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
