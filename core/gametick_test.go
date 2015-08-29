package core

import (
	"testing"

	"github.com/esanmiguelc/go-ttt-core/Godeps/_workspace/src/github.com/stretchr/testify/assert"
)

func TestItCreatesAGameStateWithCurrentMove(t *testing.T) {
	player1Type := "Human"
	player2Type := "Computer"
	movesPlayed := []int{0, 1, 2}
	gameState := GameTick(player1Type, player2Type, THREE_BY_THREE, movesPlayed)
	assert.Equal(t, []int{0, 1, 2, 4}, gameState.MovesPlayed)
}

func TestItReturnsTheFirstMove(t *testing.T) {
	player1Type := "Computer"
	player2Type := "Human"
	movesPlayed := []int{}
	gameState := GameTick(player1Type, player2Type, THREE_BY_THREE, movesPlayed)
	assert.Equal(t, []int{0}, gameState.MovesPlayed)
}

func TestItDoesNotCareForHumans(t *testing.T) {
	player1Type := "Computer"
	player2Type := "Human"
	movesPlayed := []int{0, 1, 2}
	gameState := GameTick(player1Type, player2Type, THREE_BY_THREE, movesPlayed)
	assert.Equal(t, []int{0, 1, 2}, gameState.MovesPlayed)
}

func TestGameStateHasTheRequiredFieldsToRebuildTheBoard(t *testing.T) {
	player1Type := "Computer"
	player2Type := "Human"
	board := Board{Slots: []string{"X", "O", "X", "", "", "", "", "", ""}}
	movesPlayed := []int{0, 1, 2}
	gameState := GameTick(player1Type, player2Type, THREE_BY_THREE, movesPlayed)
	assert.Equal(t, movesPlayed, gameState.MovesPlayed)
	assert.Equal(t, player1Type, gameState.PlayerOneType)
	assert.Equal(t, player2Type, gameState.PlayerTwoType)
	assert.Equal(t, THREE_BY_THREE, gameState.BoardSize)
	assert.Equal(t, THREE_BY_THREE, gameState.BoardSize)
	assert.Equal(t, board, gameState.Board)
	assert.False(t, gameState.CurrentPlayer.IsComputer())
	assert.Equal(t, board.Slots, gameState.BoardState)
}
