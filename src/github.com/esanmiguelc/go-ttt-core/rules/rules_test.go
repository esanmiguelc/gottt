package rules

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const THREE_BY_THREE = 3
const FIRST_PLAYER = "X"

func TestAllMovesAvailableWhenBoardIsEmpty(t *testing.T) {
	slots := createSlots(THREE_BY_THREE)
	assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8}, MovesAvailable(slots))
}

func TestReturnsMovesRemainingOnBoard(t *testing.T) {
	slots := createSlots(THREE_BY_THREE)
	addMarkToPositions(slots, FIRST_PLAYER, 4)
	assert.Equal(t, []int{0, 1, 2, 3, 5, 6, 7, 8}, MovesAvailable(slots))
}

func TestNoRowWinnerOnEmptyBoard(t *testing.T) {
	slots := createSlots(THREE_BY_THREE)
	assert.False(t, RowWinner(slots, FIRST_PLAYER))
}

func TestItHasAWinnerIfSameMarkInFirstRow(t *testing.T) {
	slots := createSlots(THREE_BY_THREE)
	addMarkToPositions(slots, FIRST_PLAYER, 0, 1, 2)
	assert.True(t, RowWinner(slots, FIRST_PLAYER))
}

func TestItHasAWinnerIfSameMarkInSecondRow(t *testing.T) {
	slots := createSlots(THREE_BY_THREE)
	addMarkToPositions(slots, FIRST_PLAYER, 3, 4, 5)
	assert.True(t, RowWinner(slots, FIRST_PLAYER))
}

func TestItHasAWinnerIfSameMarkInThirdRow(t *testing.T) {
	slots := createSlots(THREE_BY_THREE)
	addMarkToPositions(slots, FIRST_PLAYER, 6, 7, 8)
	assert.True(t, RowWinner(slots, FIRST_PLAYER))
}

func TestItHasNoWinnerIfOnlyOneMoveWasMade(t *testing.T) {
	slots := createSlots(THREE_BY_THREE)
	addMarkToPositions(slots, FIRST_PLAYER, 1)
	assert.False(t, ColumnWinner(slots, FIRST_PLAYER))
	assert.False(t, RowWinner(slots, FIRST_PLAYER))
}

func TestNoColumnWinnerOnEmptyBoard(t *testing.T) {
	slots := createSlots(THREE_BY_THREE)
	assert.False(t, ColumnWinner(slots, FIRST_PLAYER))
}

func TestItDoesNotHaveWinnerWhenNoCombinationsFoundForColumn(t *testing.T) {
	slots := createSlots(THREE_BY_THREE)
	addMarkToPositions(slots, FIRST_PLAYER, 3, 1, 8)
	assert.False(t, ColumnWinner(slots, FIRST_PLAYER))
}

func TestItHasAWinnerIfSameMarkInFirstColumn(t *testing.T) {
	slots := createSlots(THREE_BY_THREE)
	addMarkToPositions(slots, FIRST_PLAYER, 0, 3, 6)
	assert.True(t, ColumnWinner(slots, FIRST_PLAYER))
}

func TestItHasAWinnerIfSameMarkInSecondColumn(t *testing.T) {
	slots := createSlots(THREE_BY_THREE)
	addMarkToPositions(slots, FIRST_PLAYER, 1, 4, 7)
	assert.True(t, ColumnWinner(slots, FIRST_PLAYER))
}

func TestItHasAWinnerIfSameMarkInThirdColumn(t *testing.T) {
	slots := createSlots(THREE_BY_THREE)
	addMarkToPositions(slots, FIRST_PLAYER, 2, 5, 8)
	assert.True(t, ColumnWinner(slots, FIRST_PLAYER))
}

func TestNoDiagonalWinnerIfBoardEmpty(t *testing.T) {
	slots := createSlots(THREE_BY_THREE)
	assert.False(t, DiagonalWinner(slots, FIRST_PLAYER))
}

func TestItHasAWinnerIfSameMarkInLeftToRightDiagonal(t *testing.T) {
	slots := createSlots(THREE_BY_THREE)
	addMarkToPositions(slots, FIRST_PLAYER, 0, 4, 8)
	assert.True(t, DiagonalWinner(slots, FIRST_PLAYER))
}

func TestItHasAWinnerIfSameMarkInRightToLeftDiagonal(t *testing.T) {
	slots := createSlots(THREE_BY_THREE)
	addMarkToPositions(slots, FIRST_PLAYER, 2, 4, 6)
	assert.True(t, DiagonalWinner(slots, FIRST_PLAYER))
}

func TestItReturnsTrueWhenThereIsAWinningCombinationDiagonallyForMark(t *testing.T) {
	slots := createSlots(THREE_BY_THREE)
	addMarkToPositions(slots, FIRST_PLAYER, 0, 4, 8)
	assert.True(t, IsWinner(slots, "X"))
}

func TestItReturnsTrueWhenThereIsAWinningCombinationRowForMark(t *testing.T) {
	slots := createSlots(THREE_BY_THREE)
	addMarkToPositions(slots, FIRST_PLAYER, 0, 1, 2)
	assert.True(t, IsWinner(slots, "X"))
}

func TestItReturnsTrueWhenThereIsAWinningCombinationColumnForMark(t *testing.T) {
	slots := createSlots(THREE_BY_THREE)
	addMarkToPositions(slots, FIRST_PLAYER, 0, 3, 6)
	assert.True(t, IsWinner(slots, "X"))
}

func createSlots(size int) []string {
	return make([]string, size*size)
}

func addMarkToPositions(slots []string, mark string, positions ...int) {
	for _, position := range positions {
		slots[position] = mark
	}
}
