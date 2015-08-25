package rules

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const THREE_BY_THREE = 3
const FIRST_PLAYER = "X"

func TestAllMovesAvailableWhenBoardIsEmpty(t *testing.T) {
	slots := createSlots(THREE_BY_THREE)
	assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8}, MovesAvailable(slots), "Has all the moves available")
}

func TestReturnsMovesRemainingOnBoard(t *testing.T) {
	slots := createSlots(THREE_BY_THREE)
	addMarkToPositions(slots, FIRST_PLAYER, 4)
	assert.Equal(t, []int{0, 1, 2, 3, 5, 6, 7, 8}, MovesAvailable(slots), "Returns all the moves remaining")
}

func TestNoRowWinnerOnEmptyBoard(t *testing.T) {
	slots := createSlots(THREE_BY_THREE)
	assert.False(t, RowWinner(slots, FIRST_PLAYER), "Returns False")
}

func TestItHasAWinnerIfSameMarkInFirstRow(t *testing.T) {
	slots := createSlots(THREE_BY_THREE)
	addMarkToPositions(slots, FIRST_PLAYER, 0, 1, 2)
	assert.True(t, RowWinner(slots, FIRST_PLAYER), "Returns true")
}

func TestItHasAWinnerIfSameMarkInSecondRow(t *testing.T) {
	slots := createSlots(THREE_BY_THREE)
	addMarkToPositions(slots, FIRST_PLAYER, 3, 4, 5)
	assert.True(t, RowWinner(slots, FIRST_PLAYER), "Returns true")
}

func TestItHasAWinnerIfSameMarkInThirdRow(t *testing.T) {
	slots := createSlots(THREE_BY_THREE)
	addMarkToPositions(slots, FIRST_PLAYER, 6, 7, 8)
	assert.True(t, RowWinner(slots, FIRST_PLAYER), "Returns true")
}

func TestItHasNoWinnerIfOnlyOneMoveWasMade(t *testing.T) {
	slots := createSlots(THREE_BY_THREE)
	addMarkToPositions(slots, FIRST_PLAYER, 1)
	assert.False(t, ColumnWinner(slots, FIRST_PLAYER), "Returns false")
	assert.False(t, RowWinner(slots, FIRST_PLAYER), "Returns false")
}

func TestNoColumnWinnerOnEmptyBoard(t *testing.T) {
	slots := createSlots(THREE_BY_THREE)
	assert.False(t, ColumnWinner(slots, FIRST_PLAYER), "Returns false")
}

func TestItDoesNotHaveWinnerWhenNoCombinationsFoundForColumn(t *testing.T) {
	slots := createSlots(THREE_BY_THREE)
	addMarkToPositions(slots, FIRST_PLAYER, 3, 1, 8)
	assert.False(t, ColumnWinner(slots, FIRST_PLAYER), "Returns false")
}

func TestItHasAWinnerIfSameMarkInFirstColumn(t *testing.T) {
	slots := createSlots(THREE_BY_THREE)
	addMarkToPositions(slots, FIRST_PLAYER, 0, 3, 6)
	assert.True(t, ColumnWinner(slots, FIRST_PLAYER), "Returns true")
}

func TestItHasAWinnerIfSameMarkInSecondColumn(t *testing.T) {
	slots := createSlots(THREE_BY_THREE)
	addMarkToPositions(slots, FIRST_PLAYER, 1, 4, 7)
	assert.True(t, ColumnWinner(slots, FIRST_PLAYER), "Returns true")
}

func TestItHasAWinnerIfSameMarkInThirdColumn(t *testing.T) {
	slots := createSlots(THREE_BY_THREE)
	addMarkToPositions(slots, FIRST_PLAYER, 2, 5, 8)
	assert.True(t, ColumnWinner(slots, FIRST_PLAYER), "Returns true")
}

func TestNoDiagonalWinnerIfBoardEmpty(t *testing.T) {
	slots := createSlots(THREE_BY_THREE)
	assert.False(t, DiagonalWinner(slots, FIRST_PLAYER), "Returns false")
}

func TestItHasAWinnerIfSameMarkInLeftToRightDiagonal(t *testing.T) {
	slots := createSlots(THREE_BY_THREE)
	addMarkToPositions(slots, FIRST_PLAYER, 0, 4, 8)
	assert.True(t, DiagonalWinner(slots, FIRST_PLAYER), "Returns true")
}

func TestItHasAWinnerIfSameMarkInRightToLeftDiagonal(t *testing.T) {
	slots := createSlots(THREE_BY_THREE)
	addMarkToPositions(slots, FIRST_PLAYER, 2, 4, 6)
	assert.True(t, DiagonalWinner(slots, FIRST_PLAYER), "Returns true")
}

func TestItReturnsTrueWhenThereIsAWinningCombinationDiagonallyForMark(t *testing.T) {
	slots := createSlots(THREE_BY_THREE)
	addMarkToPositions(slots, FIRST_PLAYER, 0, 4, 8)
	assert.True(t, IsWinner(slots, "X"), "Returns true")
}

func TestItReturnsTrueWhenThereIsAWinningCombinationRowForMark(t *testing.T) {
	slots := createSlots(THREE_BY_THREE)
	addMarkToPositions(slots, FIRST_PLAYER, 0, 1, 2)
	assert.True(t, IsWinner(slots, "X"), "Returns true")
}

func TestItReturnsTrueWhenThereIsAWinningCombinationColumnForMark(t *testing.T) {
	slots := createSlots(THREE_BY_THREE)
	addMarkToPositions(slots, FIRST_PLAYER, 0, 3, 6)
	assert.True(t, IsWinner(slots, "X"), "Returns true")
}

func createSlots(size int) []string {
	return make([]string, size*size)
}

func addMarkToPositions(slots []string, mark string, positions ...int) {
	for _, position := range positions {
		slots[position] = mark
	}
}
