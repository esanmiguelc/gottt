package core

import (
	"github.com/esanmiguelc/go-ttt-core/Godeps/_workspace/src/github.com/stretchr/testify/assert"
	"testing"
)

func TestItCreatesABoardWithTheCorrectSize(t *testing.T) {
	board := NewBoard(3)
	assert.Equal(t, 9, len(board.Slots))
}

func TestItCreatesAnEmptyBoard(t *testing.T) {
	board := NewBoard(3)
	assert.Equal(t, make([]string, 9), board.Slots)
}

func TestPlacesAMoveOnBoard(t *testing.T) {
	board := NewBoard(3)
	testBoard := make([]string, 9)
	testBoard[0] = "x"
	board.PlaceMove(0, "x")
	assert.Equal(t, testBoard, board.Slots)
}

func TestBoardIsNotFull(t *testing.T) {
	board := CreateBoard(THREE_BY_THREE)
	assert.False(t, board.Full())
}

func TestBoardIsFull(t *testing.T) {
	board := CreateBoard(THREE_BY_THREE)
	AddMarkToPositions(board, FIRST_PLAYER, 0, 1, 2, 3, 4, 5, 6, 7, 8)
	assert.True(t, board.Full())
}

func TestUndoMove(t *testing.T) {
	board := CreateBoard(THREE_BY_THREE)
	AddMarkToPositions(board, FIRST_PLAYER, 0)
	board.Undo(0)
	assert.Equal(t, "", board.Slots[0])
}
