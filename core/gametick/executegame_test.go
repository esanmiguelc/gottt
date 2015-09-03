package gametick_test

import (
	"testing"

	. "github.com/esanmiguelc/gottt/core/gametick"
	"github.com/esanmiguelc/gottt/core/rules"
	"github.com/stretchr/testify/assert"
)

func TestItReturnsTheGameState(t *testing.T) {
	movesPlayed := []int{0}
	gameState, _, _ := ExecuteGame(rules.HUMAN, rules.COMPUTER, rules.THREE_BY_THREE, movesPlayed)
	assert.Equal(t, []int{0, 4}, gameState.MovesPlayed)
}
func TestItReturnsFalseIfGameIsNotOver(t *testing.T) {
	movesPlayed := []int{0}
	_, gameOver, _ := ExecuteGame(rules.HUMAN, rules.COMPUTER, rules.THREE_BY_THREE, movesPlayed)
	assert.False(t, gameOver)
}

func TestItReturnsTrueIfGameIsOver(t *testing.T) {
	movesPlayed := []int{0, 1, 3, 4, 6}
	_, gameOver, _ := ExecuteGame(rules.HUMAN, rules.COMPUTER, rules.THREE_BY_THREE, movesPlayed)
	assert.True(t, gameOver)
}

func TestItReturnsTieForResult(t *testing.T) {
	movesPlayed := []int{0, 4, 8, 1, 7, 6, 2, 5, 3}
	_, _, result := ExecuteGame(rules.HUMAN, rules.COMPUTER, rules.THREE_BY_THREE, movesPlayed)
	assert.Equal(t, rules.TIE, result)
}

func TestItReturnsMarkFirstPlayerForResult(t *testing.T) {
	movesPlayed := []int{0, 1, 3, 4, 6}
	_, _, result := ExecuteGame(rules.HUMAN, rules.COMPUTER, rules.THREE_BY_THREE, movesPlayed)
	assert.Equal(t, rules.FIRST_PLAYER, result)
}

func TestItReturnsMarkSecondPlayerForResult(t *testing.T) {
	movesPlayed := []int{0, 1, 2, 4, 3, 7}
	_, _, result := ExecuteGame(rules.HUMAN, rules.COMPUTER, rules.THREE_BY_THREE, movesPlayed)
	assert.Equal(t, rules.SECOND_PLAYER, result)
}
