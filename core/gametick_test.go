package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItCreatesAGameStateWithCurrentMove(t *testing.T) {
	movesPlayed := []int{0, 1, 2}
	gameState := GameTick(HUMAN, COMPUTER, THREE_BY_THREE, movesPlayed)
	assert.Equal(t, []int{0, 1, 2, 4}, gameState.MovesPlayed)
}

func TestItReturnsTheFirstMove(t *testing.T) {
	movesPlayed := []int{}
	gameState := GameTick(COMPUTER, HUMAN, THREE_BY_THREE, movesPlayed)
	assert.Equal(t, []int{0}, gameState.MovesPlayed)
}

func TestItDoesNotCareForHumans(t *testing.T) {
	movesPlayed := []int{0, 1, 2}
	gameState := GameTick(COMPUTER, HUMAN, THREE_BY_THREE, movesPlayed)
	assert.Equal(t, []int{0, 1, 2}, gameState.MovesPlayed)
}

func TestGameStateHasTheRequiredFieldsToRebuildTheBoard(t *testing.T) {
	board := Board{Slots: []string{"X", "O", "X", "", "", "", "", "", ""}}
	movesPlayed := []int{0, 1, 2}
	gameState := GameTick(COMPUTER, HUMAN, THREE_BY_THREE, movesPlayed)
	assert.Equal(t, movesPlayed, gameState.MovesPlayed)
	assert.Equal(t, COMPUTER, gameState.PlayerOneType)
	assert.Equal(t, HUMAN, gameState.PlayerTwoType)
	assert.Equal(t, THREE_BY_THREE, gameState.BoardSize)
	assert.Equal(t, THREE_BY_THREE, gameState.BoardSize)
	assert.Equal(t, board, gameState.Board)
	assert.False(t, gameState.CurrentPlayer.IsComputer())
	assert.Equal(t, board.Slots, gameState.BoardState)
}
