package tttboard

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestItCreatesABoardWithTheCorrectSize(t *testing.T) {
	board := createBoard(3)
	assert.Equal(t, 9, len(board.Slots), "This is truthy")
}

func TestItCreatesAnEmptyBoard(t *testing.T) {
	board := createBoard(3)
	assert.Equal(t, make([]string, 9), board.Slots, "The board is empty")
}

func TestPlacesAMoveOnBoard(t *testing.T) {
	board := createBoard(3)
	testBoard := make([]string, 9)
	testBoard[0] = "x"
	board.PlaceMove(0, "x")
	assert.Equal(t, testBoard, board.Slots, "Other test")
}
