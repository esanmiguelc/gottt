package gamestate

import (
	"github.com/esanmiguelc/gottt/core/board"
	"github.com/esanmiguelc/gottt/core/iplayer"
	"github.com/esanmiguelc/gottt/core/rules"
)

type GameState struct {
	MovesPlayed   []int
	PlayerOneType string
	PlayerTwoType string
	CurrentPlayer iplayer.Player
	BoardSize     int
	BoardState    []string
	Board         board.Board
}

func GetCurrentPlayer(board board.Board, playerOne, playerTwo iplayer.Player) iplayer.Player {
	if IsCurrentPlayer(board, playerOne.GetMark()) {
		return playerOne
	}
	return playerTwo
}

func IsCurrentPlayer(board board.Board, mark string) bool {
	return GetCurrentMark(board) == mark
}

func GetCurrentMark(board board.Board) string {
	movesAvailableCount := len(rules.MovesAvailable(board))
	if rules.GetBoardSize(board.Slots)%3 == 0 {
		return currentMark(movesAvailableCount, 2)
	}
	return currentMark(movesAvailableCount, 3)
}

func currentMark(moveCount, num int) string {
	if moveCount%num == 0 {
		return rules.SECOND_PLAYER
	}
	return rules.FIRST_PLAYER
}

func GetOpponentMark(mark string) string {
	if mark == rules.FIRST_PLAYER {
		return rules.SECOND_PLAYER
	}
	return rules.FIRST_PLAYER
}
