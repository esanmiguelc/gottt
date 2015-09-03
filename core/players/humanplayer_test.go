package players_test

import (
	"testing"

	"github.com/esanmiguelc/gottt/core/board"
	. "github.com/esanmiguelc/gottt/core/players"
	"github.com/esanmiguelc/gottt/core/rules"
	"github.com/stretchr/testify/assert"
)

func TestAlwaysReturnsZero(t *testing.T) {
	board := board.NewBoard(rules.THREE_BY_THREE)
	player := NewHumanPlayer(rules.FIRST_PLAYER)
	assert.Equal(t, 0, player.GetMove(board, rules.FIRST_PLAYER, rules.SECOND_PLAYER))
}

func TestIsNotAComputer(t *testing.T) {
	player := NewHumanPlayer(rules.FIRST_PLAYER)
	assert.False(t, player.IsComputer())
}

func TestGetsTheMark(t *testing.T) {
	player := NewHumanPlayer(rules.FIRST_PLAYER)
	assert.Equal(t, rules.FIRST_PLAYER, player.GetMark())
}
