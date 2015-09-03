package validators_test

import (
	"testing"

	"github.com/esanmiguelc/gottt/core/rules"
	. "github.com/esanmiguelc/gottt/web/validators"
	"github.com/stretchr/testify/assert"
)

func TestItReturnsErrorIfBoardIsNotThree(t *testing.T) {
	_, _, _, _, err := ValidateParams(rules.HUMAN, rules.COMPUTER, "4", "0")
	assert.NotNil(t, err)
}

func TestItReturnsErrorIfInputForThreeByThreeIsNine(t *testing.T) {
	_, _, _, _, err := ValidateParams(rules.HUMAN, rules.COMPUTER, "3", "9")
	assert.NotNil(t, err)
}

func TestItReturnsErrorIfItCantParseBoardSize(t *testing.T) {
	_, _, _, _, err := ValidateParams(rules.HUMAN, rules.COMPUTER, "T", "0")
	assert.NotNil(t, err)
}

func TestItReturnsErrorIfItIsNotHumanOrComputerPlayerOne(t *testing.T) {
	_, _, _, _, err := ValidateParams("Hum", rules.COMPUTER, "3", "0")
	assert.NotNil(t, err)
}

func TestItReturnsErrorIfItIsNotHumanOrComputerPlayerTwo(t *testing.T) {
	_, _, _, _, err := ValidateParams(rules.HUMAN, "Comput", "3", "0")
	assert.NotNil(t, err)
}

func TestItReturnsBoardSizeAsInt(t *testing.T) {
	_, _, boardSize, _, _ := ValidateParams(rules.HUMAN, rules.COMPUTER, "3", "0")
	assert.Equal(t, 3, boardSize)
}

func TestItReturnsMovesPlayedAsArray(t *testing.T) {
	_, _, _, movesPlayed, _ := ValidateParams(rules.HUMAN, rules.COMPUTER, "3", "0")
	assert.Equal(t, []int{0}, movesPlayed)
}

func TestItLeavesPlayersAsIs(t *testing.T) {
	firstPlayer, secondPlayer, _, _, _ := ValidateParams(rules.HUMAN, rules.COMPUTER, "3", "0")
	assert.Equal(t, firstPlayer, rules.HUMAN)
	assert.Equal(t, secondPlayer, rules.COMPUTER)
}

func TestWhenGivenDuplicateValuesOnlyTakesTheFirstOne(t *testing.T) {
	_, _, _, movesPlayed, _ := ValidateParams(rules.HUMAN, rules.COMPUTER, "3", "01224347")
	assert.Equal(t, []int{0, 1, 2, 4, 3, 7}, movesPlayed)
}
