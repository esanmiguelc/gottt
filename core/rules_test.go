package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
	assert.True(t, IsWinner(board, FIRST_PLAYER))
}

func TestItReturnsTrueWhenThereIsAWinningCombinationRowForMark(t *testing.T) {
	board := CreateBoard(THREE_BY_THREE)
	AddMarkToPositions(board, FIRST_PLAYER, 0, 1, 2)
	assert.True(t, IsWinner(board, FIRST_PLAYER))
}

func TestItReturnsTrueWhenThereIsAWinningCombinationColumnForMark(t *testing.T) {
	board := CreateBoard(THREE_BY_THREE)
	AddMarkToPositions(board, FIRST_PLAYER, 0, 3, 6)
	assert.True(t, IsWinner(board, FIRST_PLAYER))
}

func TestGameIsOverWhenBoardIsFull(t *testing.T) {
	board := CreateBoard(THREE_BY_THREE)
	AddMarkToPositions(board, FIRST_PLAYER, 0, 1, 2, 3, 4, 5, 6, 7, 8)
	assert.True(t, IsGameOver(board))
}

func TestGameIsOverWhenFirstPlayerWins(t *testing.T) {
	board := CreateBoard(THREE_BY_THREE)
	AddMarkToPositions(board, FIRST_PLAYER, 0, 1, 2)
	assert.True(t, IsGameOver(board))
}

func TestGameIsOverWhenSecondPlayerWins(t *testing.T) {
	board := CreateBoard(THREE_BY_THREE)
	AddMarkToPositions(board, SECOND_PLAYER, 0, 1, 2)
	assert.True(t, IsGameOver(board))
}

func TestCurrentPlayerIsFirstPlayerForFirstMove(t *testing.T) {
	board := CreateBoard(THREE_BY_THREE)
	assert.Equal(t, FIRST_PLAYER, GetCurrentMark(board))
}

func TestCurrentPlayerIsSecondPlayer(t *testing.T) {
	board := CreateBoard(THREE_BY_THREE)
	AddMarkToPositions(board, FIRST_PLAYER, 0)
	assert.Equal(t, SECOND_PLAYER, GetCurrentMark(board))
}

func TestCurrentPlayerIsFirstPlayerAfterTwoMoves(t *testing.T) {
	board := CreateBoard(THREE_BY_THREE)
	AddMarkToPositions(board, FIRST_PLAYER, 0)
	AddMarkToPositions(board, SECOND_PLAYER, 1)
	assert.Equal(t, FIRST_PLAYER, GetCurrentMark(board))
}

func TestCurrentPlayerIsFirstPlayerForFirstMoveOnFourByFour(t *testing.T) {
	board := CreateBoard(FOUR_BY_FOUR)
	AddMarkToPositions(board, FIRST_PLAYER, 0)
	assert.Equal(t, SECOND_PLAYER, GetCurrentMark(board))
}

func TestCurrentPlayerIsSecondPlayerOnFourByFour(t *testing.T) {
	board := CreateBoard(FOUR_BY_FOUR)
	assert.Equal(t, FIRST_PLAYER, GetCurrentMark(board))
}

func TestIsNotCurrentPlayer(t *testing.T) {
	board := CreateBoard(FOUR_BY_FOUR)
	assert.False(t, IsCurrentPlayer(board, SECOND_PLAYER))
}

func TestIsCurrentPlayer(t *testing.T) {
	board := CreateBoard(FOUR_BY_FOUR)
	assert.True(t, IsCurrentPlayer(board, FIRST_PLAYER))
}

func TestGetsWinnerMark(t *testing.T) {
	board := CreateBoard(THREE_BY_THREE)
	AddMarkToPositions(board, FIRST_PLAYER, 0, 1, 2)
	assert.Equal(t, FIRST_PLAYER, GetResult(board))
}

func TestGetsWinnerMarkSecondPlayer(t *testing.T) {
	board := CreateBoard(THREE_BY_THREE)
	AddMarkToPositions(board, SECOND_PLAYER, 0, 1, 2)
	assert.Equal(t, SECOND_PLAYER, GetResult(board))
}

func TestGetsTieIfNoWinner(t *testing.T) {
	board := CreateBoard(THREE_BY_THREE)
	AddMarkToPositions(board, FIRST_PLAYER, 0, 1, 5, 6, 8)
	AddMarkToPositions(board, SECOND_PLAYER, 2, 3, 4, 7)
	assert.Equal(t, "Tie", GetResult(board))
}

func TestGetCurrentPlayer(t *testing.T) {
	board := CreateBoard(THREE_BY_THREE)
	playerOne := HumanPlayer{Mark: FIRST_PLAYER}
	playerTwo := HumanPlayer{Mark: SECOND_PLAYER}
	assert.Equal(t, playerOne, GetCurrentPlayer(board, playerOne, playerTwo))
}
