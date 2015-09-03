package gametick

import (
	"github.com/esanmiguelc/gottt/core/gamestate"
	"github.com/esanmiguelc/gottt/core/rules"
)

func ExecuteGame(playerOneType, playerTwoType string, boardSize int, movesPlayed []int) (gamestate.GameState, bool, string) {
	gameState := GameTick(playerOneType, playerTwoType, boardSize, movesPlayed)

	return gameState, rules.IsGameOver(gameState.Board), rules.GetResult(gameState.Board)
}
