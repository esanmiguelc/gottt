package gametick_test

import (
	"testing"

	. "github.com/esanmiguelc/gottt/core/gametick"
	"github.com/esanmiguelc/gottt/core/rules"
	"github.com/stretchr/testify/assert"
)

func TestItReturnsTheGameState(t *testing.T) {
	gameState, _, _, _ := ExecuteGame(rules.HUMAN, rules.COMPUTER, "3", "0")
	assert.Equal(t, []int{0, 4}, gameState.MovesPlayed)
}

func TestItReturnsErrorIfBoardIsNotThree(t *testing.T) {
	_, _, _, err := ExecuteGame(rules.HUMAN, rules.COMPUTER, "4", "0")
	assert.NotNil(t, err)
}

func TestItReturnsErrorIfInputForThreeByThreeIsNine(t *testing.T) {
	_, _, _, err := ExecuteGame(rules.HUMAN, rules.COMPUTER, "3", "9")
	assert.NotNil(t, err)
}

func TestItReturnsErrorIfItCantParseBoardSize(t *testing.T) {
	_, _, _, err := ExecuteGame(rules.HUMAN, rules.COMPUTER, "T", "0")
	assert.NotNil(t, err)
}

func TestItReturnsErrorIfItIsNotHumanOrComputerPlayerOne(t *testing.T) {
	_, _, _, err := ExecuteGame("Hum", rules.COMPUTER, "3", "0")
	assert.NotNil(t, err)
}

func TestItReturnsErrorIfItIsNotHumanOrComputerPlayerTwo(t *testing.T) {
	_, _, _, err := ExecuteGame(rules.HUMAN, "Comput", "3", "0")
	assert.NotNil(t, err)
}

func TestItReturnsFalseIfGameIsNotOver(t *testing.T) {
	_, gameOver, _, _ := ExecuteGame(rules.HUMAN, rules.COMPUTER, "3", "0")
	assert.False(t, gameOver)
}

func TestItReturnsTrueIfGameIsOver(t *testing.T) {
	_, gameOver, _, _ := ExecuteGame(rules.HUMAN, rules.COMPUTER, "3", "01346")
	assert.True(t, gameOver)
}

func TestItReturnsTieForResult(t *testing.T) {
	_, _, result, _ := ExecuteGame(rules.HUMAN, rules.COMPUTER, "3", "048176253")
	assert.Equal(t, rules.TIE, result)
}

func TestItReturnsMarkFirstPlayerForResult(t *testing.T) {
	_, _, result, _ := ExecuteGame(rules.HUMAN, rules.COMPUTER, "3", "01346")
	assert.Equal(t, rules.FIRST_PLAYER, result)
}

func TestItReturnsMarkSecondPlayerForResult(t *testing.T) {
	_, _, result, _ := ExecuteGame(rules.HUMAN, rules.COMPUTER, "3", "012437")
	assert.Equal(t, rules.SECOND_PLAYER, result)
}

func TestWhenGivenDuplicateValuesOnlyTakesTheFirstOne(t *testing.T) {
	_, _, result, _ := ExecuteGame(rules.HUMAN, rules.COMPUTER, "3", "01224347")
	assert.Equal(t, rules.SECOND_PLAYER, result)
}
