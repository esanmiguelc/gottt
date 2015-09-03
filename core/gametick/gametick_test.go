package gametick_test

import (
	"testing"

	"github.com/esanmiguelc/gottt/core/board"
	. "github.com/esanmiguelc/gottt/core/gametick"
	"github.com/esanmiguelc/gottt/core/rules"
	"github.com/stretchr/testify/assert"
)

func TestItCreatesAGameStateWithCurrentMove(t *testing.T) {
	movesPlayed := []int{0, 1, 2}
	gameState := GameTick(rules.HUMAN, rules.COMPUTER, rules.THREE_BY_THREE, movesPlayed)
	assert.Equal(t, []int{0, 1, 2, 4}, gameState.MovesPlayed)
}

func TestItReturnsTheFirstMove(t *testing.T) {
	movesPlayed := []int{}
	gameState := GameTick(rules.COMPUTER, rules.HUMAN, rules.THREE_BY_THREE, movesPlayed)
	assert.Equal(t, []int{0}, gameState.MovesPlayed)
}

func TestItDoesNotCareForHumans(t *testing.T) {
	movesPlayed := []int{0, 1, 2}
	gameState := GameTick(rules.COMPUTER, rules.HUMAN, rules.THREE_BY_THREE, movesPlayed)
	assert.Equal(t, []int{0, 1, 2}, gameState.MovesPlayed)
}

func TestGameStateHasTheRequiredFieldsToRebuildTheBoard(t *testing.T) {
	board := board.Board{Slots: []string{rules.FIRST_PLAYER, rules.SECOND_PLAYER, rules.FIRST_PLAYER, "", "", "", "", "", ""}}
	movesPlayed := []int{0, 1, 2}
	gameState := GameTick(rules.COMPUTER, rules.HUMAN, rules.THREE_BY_THREE, movesPlayed)
	assert.Equal(t, movesPlayed, gameState.MovesPlayed)
	assert.Equal(t, rules.COMPUTER, gameState.PlayerOneType)
	assert.Equal(t, rules.HUMAN, gameState.PlayerTwoType)
	assert.Equal(t, rules.THREE_BY_THREE, gameState.BoardSize)
	assert.Equal(t, rules.THREE_BY_THREE, gameState.BoardSize)
	assert.Equal(t, board, gameState.Board)
	assert.False(t, gameState.CurrentPlayer.IsComputer())
	assert.Equal(t, board.Slots, gameState.BoardState)
}
