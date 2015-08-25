package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestItCreatesABoardWithTheCorrectSize(t *testing.T) {
	board := createBoard(3)
	assert.Equal(t, 9, len(board.Slots))
}

func TestItCreatesAnEmptyBoard(t *testing.T) {
	board := createBoard(3)
	assert.Equal(t, make([]string, 9), board.Slots)
}

func TestPlacesAMoveOnBoard(t *testing.T) {
	board := createBoard(3)
	testBoard := make([]string, 9)
	testBoard[0] = "x"
	board.PlaceMove(0, "x")
	assert.Equal(t, testBoard, board.Slots)
}
