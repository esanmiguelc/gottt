package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlwaysReturnsZero(t *testing.T) {
	board := NewBoard(3)
	player := NewHumanPlayer(FIRST_PLAYER)
	assert.Equal(t, 0, player.GetMove(board, FIRST_PLAYER, SECOND_PLAYER))
}

func TestIsNotAComputer(t *testing.T) {
	player := NewHumanPlayer(FIRST_PLAYER)
	assert.False(t, player.IsComputer())
}

func TestGetsTheMark(t *testing.T) {
	player := NewHumanPlayer(FIRST_PLAYER)
	assert.Equal(t, FIRST_PLAYER, player.GetMark())
}
