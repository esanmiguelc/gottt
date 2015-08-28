package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItCreatesAGameStateWithCurrentMove(t *testing.T) {
	player1Type := "Human"
	player2Type := "Computer"
	movesPlayed := []int{0, 1, 2}
	gameState := GameTick(player1Type, player2Type, THREE_BY_THREE, movesPlayed)
	assert.Equal(t, []int{0, 1, 2, 4}, gameState.MovesPlayed)
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
	movesPlayed := []int{0, 1, 2}
	gameState := GameTick(player1Type, player2Type, THREE_BY_THREE, movesPlayed)
	assert.Equal(t, movesPlayed, gameState.MovesPlayed)
	assert.Equal(t, player1Type, gameState.PlayerOneType)
	assert.Equal(t, player2Type, gameState.PlayerTwoType)
	assert.Equal(t, THREE_BY_THREE, gameState.BoardSize)
}
