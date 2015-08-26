package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestScoreWhenNeitherPlayerWins(t *testing.T) {
	board := CreateBoard(3)
	assert.Equal(t, 0, Score(board, FIRST_PLAYER, SECOND_PLAYER))
}

func TestScoreWhenIWin(t *testing.T) {
	board := CreateBoard(3)
	AddMarkToPositions(board, FIRST_PLAYER, 0, 1, 2)
	assert.Equal(t, 10, Score(board, FIRST_PLAYER, SECOND_PLAYER))
}

func TestScoreWhenOpponentWins(t *testing.T) {
	board := CreateBoard(3)
	AddMarkToPositions(board, SECOND_PLAYER, 0, 1, 2)
	assert.Equal(t, -10, Score(board, FIRST_PLAYER, SECOND_PLAYER))
}

func TestGetsTheFirstMove(t *testing.T) {
	board := CreateBoard(3)
	assert.Equal(t, 0, GetMove(board, FIRST_PLAYER, SECOND_PLAYER))
}

func TestGetsTheCenterMove(t *testing.T) {
	board := CreateBoard(3)
	AddMarkToPositions(board, FIRST_PLAYER, 0)
	assert.Equal(t, 4, GetMove(board, SECOND_PLAYER, FIRST_PLAYER))
}

func TestBlocksAWin(t *testing.T) {
	board := CreateBoard(3)
	AddMarkToPositions(board, FIRST_PLAYER, 0, 2, 5, 7)
	AddMarkToPositions(board, SECOND_PLAYER, 1, 4, 3)
	assert.Equal(t, 8, GetMove(board, SECOND_PLAYER, FIRST_PLAYER))
}

func TestGoesForTheWin(t *testing.T) {
	board := CreateBoard(3)
	AddMarkToPositions(board, FIRST_PLAYER, 0, 5, 2)
	AddMarkToPositions(board, SECOND_PLAYER, 1, 4)
	assert.Equal(t, 7, GetMove(board, SECOND_PLAYER, FIRST_PLAYER))
}
