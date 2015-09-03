package gametick

import (
	"github.com/esanmiguelc/gottt/core/board"
	"github.com/esanmiguelc/gottt/core/gamestate"
	"github.com/esanmiguelc/gottt/core/iplayer"
	"github.com/esanmiguelc/gottt/core/players"
	"github.com/esanmiguelc/gottt/core/rules"
)

func GameTick(playerOneType, playerTwoType string, boardSize int, movesPlayed []int) gamestate.GameState {
	playerOne := createFirstPlayer(playerOneType)
	playerTwo := createSecondPlayer(playerTwoType)
	board := board.NewBoard(boardSize)
	rebuildBoard(board, movesPlayed)

	return gamestate.GameState{
		MovesPlayed:   takeTurn(board, playerOne, playerTwo, movesPlayed),
		PlayerOneType: playerOneType,
		PlayerTwoType: playerTwoType,
		CurrentPlayer: gamestate.GetCurrentPlayer(board, playerOne, playerTwo),
		BoardSize:     boardSize,
		BoardState:    board.Slots,
		Board:         board,
	}
}

func rebuildBoard(board board.Board, movesPlayed []int) {
	for _, move := range movesPlayed {
		board.PlaceMove(move, gamestate.GetCurrentMark(board))
	}
}

func takeTurn(board board.Board, playerOne, playerTwo iplayer.Player, movesPlayed []int) []int {
	currentPlayer := gamestate.GetCurrentPlayer(board, playerOne, playerTwo)
	if currentPlayer.IsComputer() {
		move := currentPlayer.GetMove(board, currentPlayer.GetMark(), rules.GetOpponentMark(currentPlayer.GetMark()))
		board.PlaceMove(move, currentPlayer.GetMark())
		movesPlayed = append(movesPlayed, move)
	}
	return movesPlayed
}

func createFirstPlayer(playerType string) iplayer.Player {
	return createPlayer(playerType, rules.FIRST_PLAYER)
}

func createSecondPlayer(playerType string) iplayer.Player {
	return createPlayer(playerType, rules.SECOND_PLAYER)
}

func createPlayer(playerType, mark string) iplayer.Player {
	if playerType == rules.COMPUTER {
		return players.ImpossiblePlayer{Mark: mark}
	}
	return players.HumanPlayer{mark}
}
