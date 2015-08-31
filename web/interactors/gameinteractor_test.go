package interactors

import (
	"testing"

	"github.com/esanmiguelc/gottt/core"
	"github.com/stretchr/testify/assert"
)

func TestItReturnsTheGameState(t *testing.T) {
	gameState, _, _, _ := ExecuteGameInteractor(core.HUMAN, core.COMPUTER, "3", "0")
	assert.Equal(t, []int{0, 4}, gameState.MovesPlayed)
}

func TestItReturnsErrorIfBoardIsNotThree(t *testing.T) {
	_, _, _, err := ExecuteGameInteractor(core.HUMAN, core.COMPUTER, "4", "0")
	assert.NotNil(t, err)
}

func TestItReturnsErrorIfItCantParseBoardSize(t *testing.T) {
	_, _, _, err := ExecuteGameInteractor(core.HUMAN, core.COMPUTER, "T", "0")
	assert.NotNil(t, err)
}

func TestItReturnsErrorIfItIsNotHumanOrComputerPlayerOne(t *testing.T) {
	_, _, _, err := ExecuteGameInteractor("Hum", core.COMPUTER, "3", "0")
	assert.NotNil(t, err)
}

func TestItReturnsErrorIfItIsNotHumanOrComputerPlayerTwo(t *testing.T) {
	_, _, _, err := ExecuteGameInteractor(core.HUMAN, "Comput", "3", "0")
	assert.NotNil(t, err)
}

func TestItReturnsFalseIfGameIsNotOver(t *testing.T) {
	_, gameOver, _, _ := ExecuteGameInteractor(core.HUMAN, core.COMPUTER, "3", "0")
	assert.False(t, gameOver)
}

func TestItReturnsTrueIfGameIsOver(t *testing.T) {
	_, gameOver, _, _ := ExecuteGameInteractor(core.HUMAN, core.COMPUTER, "3", "01346")
	assert.True(t, gameOver)
}

func TestItReturnsTieForResult(t *testing.T) {
	_, _, result, _ := ExecuteGameInteractor(core.HUMAN, core.COMPUTER, "3", "048176253")
	assert.Equal(t, "Tie", result)
}

func TestItReturnsMarkFirstPlayerForResult(t *testing.T) {
	_, _, result, _ := ExecuteGameInteractor(core.HUMAN, core.COMPUTER, "3", "01346")
	assert.Equal(t, core.FIRST_PLAYER, result)
}

func TestItReturnsMarkSecondPlayerForResult(t *testing.T) {
	_, _, result, _ := ExecuteGameInteractor(core.HUMAN, core.COMPUTER, "3", "012437")
	assert.Equal(t, core.SECOND_PLAYER, result)
}
