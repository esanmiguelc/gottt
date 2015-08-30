package interactors

import (
	"testing"

	"github.com/esanmiguelc/gottt/core"
	"github.com/stretchr/testify/assert"
)

func TestItReturnsTheGameState(t *testing.T) {
	gameState, _ := ExecuteGameInteractor(core.HUMAN, core.COMPUTER, "3", "0")
	assert.Equal(t, []int{0, 4}, gameState.MovesPlayed)
}

func TestItReturnsErrorIfBoardIsNotThree(t *testing.T) {
	_, err := ExecuteGameInteractor(core.HUMAN, core.COMPUTER, "4", "0")
	assert.NotNil(t, err)
}

func TestItReturnsErrorIfItCantParseBoardSize(t *testing.T) {
	_, err := ExecuteGameInteractor(core.HUMAN, core.COMPUTER, "T", "0")
	assert.NotNil(t, err)
}

func TestItReturnsErrorIfItIsNotHumanOrComputerPlayerOne(t *testing.T) {
	_, err := ExecuteGameInteractor("Hum", core.COMPUTER, "3", "0")
	assert.NotNil(t, err)
}

func TestItReturnsErrorIfItIsNotHumanOrComputerPlayerTwo(t *testing.T) {
	_, err := ExecuteGameInteractor(core.HUMAN, "Comput", "3", "0")
	assert.NotNil(t, err)
}
