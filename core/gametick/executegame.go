package gametick

import (
	"github.com/esanmiguelc/gottt/core/gamestate"
	"github.com/esanmiguelc/gottt/core/rules"
)

func ExecuteGame(playerOneType, playerTwoType string, boardSize int, movesPlayed []int) (gamestate.GameState, bool, string) {
	movesPlayedAsArray := removeDuplicates(movesPlayed)
	gameState := GameTick(playerOneType, playerTwoType, boardSize, movesPlayedAsArray)

	return gameState, rules.IsGameOver(gameState.Board), rules.GetResult(gameState.Board)
}
