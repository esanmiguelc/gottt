package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAllMovesAvailableWhenBoardIsEmpty(t *testing.T) {
	board := NewBoard(3)
	assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8}, MovesAvailable(board))
}

func TestReturnsMovesRemainingOnBoard(t *testing.T) {
	board := NewBoard(3)
	board.PlaceMove(4, FIRST_PLAYER)
	assert.Equal(t, []int{0, 1, 2, 3, 5, 6, 7, 8}, MovesAvailable(board))
}

func TestNoRowWinnerOnEmptyBoard(t *testing.T) {
	board := CreateBoard(THREE_BY_THREE)
	assert.False(t, RowWinner(board, FIRST_PLAYER))
}

func TestItHasAWinnerIfSameMarkInFirstRow(t *testing.T) {
	board := CreateBoard(THREE_BY_THREE)
	AddMarkToPositions(board, FIRST_PLAYER, 0, 1, 2)
	assert.True(t, RowWinner(board, FIRST_PLAYER))
}

func TestItHasAWinnerIfSameMarkInSecondRow(t *testing.T) {
	board := CreateBoard(THREE_BY_THREE)
	AddMarkToPositions(board, FIRST_PLAYER, 3, 4, 5)
	assert.True(t, RowWinner(board, FIRST_PLAYER))
}

func TestItHasAWinnerIfSameMarkInThirdRow(t *testing.T) {
	board := CreateBoard(THREE_BY_THREE)
	AddMarkToPositions(board, FIRST_PLAYER, 6, 7, 8)
	assert.True(t, RowWinner(board, FIRST_PLAYER))
}

func TestItHasNoWinnerIfOnlyOneMoveWasMade(t *testing.T) {
	board := CreateBoard(THREE_BY_THREE)
	AddMarkToPositions(board, FIRST_PLAYER, 1)
	assert.False(t, ColumnWinner(board, FIRST_PLAYER))
	assert.False(t, RowWinner(board, FIRST_PLAYER))
}

func TestNoColumnWinnerOnEmptyBoard(t *testing.T) {
	board := CreateBoard(THREE_BY_THREE)
	assert.False(t, ColumnWinner(board, FIRST_PLAYER))
}

func TestItDoesNotHaveWinnerWhenNoCombinationsFoundForColumn(t *testing.T) {
	board := CreateBoard(THREE_BY_THREE)
	AddMarkToPositions(board, FIRST_PLAYER, 3, 1, 8)
	assert.False(t, ColumnWinner(board, FIRST_PLAYER))
}

func TestItHasAWinnerIfSameMarkInFirstColumn(t *testing.T) {
	board := CreateBoard(THREE_BY_THREE)
	AddMarkToPositions(board, FIRST_PLAYER, 0, 3, 6)
	assert.True(t, ColumnWinner(board, FIRST_PLAYER))
}

func TestItHasAWinnerIfSameMarkInSecondColumn(t *testing.T) {
	board := CreateBoard(THREE_BY_THREE)
	AddMarkToPositions(board, FIRST_PLAYER, 1, 4, 7)
	assert.True(t, ColumnWinner(board, FIRST_PLAYER))
}

func TestItHasAWinnerIfSameMarkInThirdColumn(t *testing.T) {
	board := CreateBoard(THREE_BY_THREE)
	AddMarkToPositions(board, FIRST_PLAYER, 2, 5, 8)
	assert.True(t, ColumnWinner(board, FIRST_PLAYER))
}

func TestNoDiagonalWinnerIfBoardEmpty(t *testing.T) {
	board := CreateBoard(THREE_BY_THREE)
	assert.False(t, DiagonalWinner(board, FIRST_PLAYER))
}

func TestItHasAWinnerIfSameMarkInLeftToRightDiagonal(t *testing.T) {
	board := CreateBoard(THREE_BY_THREE)
	AddMarkToPositions(board, FIRST_PLAYER, 0, 4, 8)
	assert.True(t, DiagonalWinner(board, FIRST_PLAYER))
}

func TestItHasAWinnerIfSameMarkInRightToLeftDiagonal(t *testing.T) {
	board := CreateBoard(THREE_BY_THREE)
	AddMarkToPositions(board, FIRST_PLAYER, 2, 4, 6)
	assert.True(t, DiagonalWinner(board, FIRST_PLAYER))
}

func TestItReturnsTrueWhenThereIsAWinningCombinationDiagonallyForMark(t *testing.T) {
	board := CreateBoard(THREE_BY_THREE)
	AddMarkToPositions(board, FIRST_PLAYER, 0, 4, 8)
	assert.True(t, IsWinner(board, "X"))
}

func TestItReturnsTrueWhenThereIsAWinningCombinationRowForMark(t *testing.T) {
	board := CreateBoard(THREE_BY_THREE)
	AddMarkToPositions(board, FIRST_PLAYER, 0, 1, 2)
	assert.True(t, IsWinner(board, "X"))
}

func TestItReturnsTrueWhenThereIsAWinningCombinationColumnForMark(t *testing.T) {
	board := CreateBoard(THREE_BY_THREE)
	AddMarkToPositions(board, FIRST_PLAYER, 0, 3, 6)
	assert.True(t, IsWinner(board, "X"))
}