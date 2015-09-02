package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItCreatesABoardWithTheCorrectSize(t *testing.T) {
	board := NewBoard(THREE_BY_THREE)
	assert.Equal(t, 9, len(board.Slots))
}

func TestItCreatesAnEmptyBoard(t *testing.T) {
	board := NewBoard(THREE_BY_THREE)
	assert.Equal(t, make([]string, 9), board.Slots)
}

func TestPlacesAMoveOnBoard(t *testing.T) {
	board := NewBoard(THREE_BY_THREE)
	testBoard := make([]string, 9)
	testBoard[0] = "x"
	board.PlaceMove(0, "x")
	assert.Equal(t, testBoard, board.Slots)
}

func TestBoardIsNotFull(t *testing.T) {
	board := NewBoard(THREE_BY_THREE)
	assert.False(t, board.Full())
}

func TestBoardIsFull(t *testing.T) {
	board := NewBoard(THREE_BY_THREE)
	AddMarkToPositions(board, FIRST_PLAYER, 0, 1, 2, 3, 4, 5, 6, 7, 8)
	assert.True(t, board.Full())
}

func TestUndoMove(t *testing.T) {
	board := NewBoard(THREE_BY_THREE)
	AddMarkToPositions(board, FIRST_PLAYER, 0)
	board.Undo(0)
	assert.Equal(t, "", board.Slots[0])
}
