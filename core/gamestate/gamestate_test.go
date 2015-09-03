package gamestate_test

import (
	"testing"

	"github.com/esanmiguelc/gottt/core/board"
	. "github.com/esanmiguelc/gottt/core/gamestate"
	"github.com/esanmiguelc/gottt/core/players"
	"github.com/esanmiguelc/gottt/core/rules"
	"github.com/esanmiguelc/gottt/core/testhelper"
	"github.com/stretchr/testify/assert"
)

func TestGetCurrentPlayer(t *testing.T) {
	board := board.NewBoard(rules.THREE_BY_THREE)
	playerOne := players.HumanPlayer{Mark: rules.FIRST_PLAYER}
	playerTwo := players.HumanPlayer{Mark: rules.SECOND_PLAYER}
	assert.Equal(t, playerOne, GetCurrentPlayer(board, playerOne, playerTwo))
}

func TestIsNotCurrentPlayer(t *testing.T) {
	board := board.NewBoard(rules.FOUR_BY_FOUR)
	assert.False(t, IsCurrentPlayer(board, rules.SECOND_PLAYER))
}

func TestIsCurrentPlayer(t *testing.T) {
	board := board.NewBoard(rules.FOUR_BY_FOUR)
	assert.True(t, IsCurrentPlayer(board, rules.FIRST_PLAYER))
}

func TestCurrentPlayerIsFirstPlayerForFirstMove(t *testing.T) {
	board := board.NewBoard(rules.THREE_BY_THREE)
	assert.Equal(t, rules.FIRST_PLAYER, GetCurrentMark(board))
}

func TestCurrentPlayerIsSecondPlayer(t *testing.T) {
	board := board.NewBoard(rules.THREE_BY_THREE)
	testhelper.AddMarkToPositions(board, rules.FIRST_PLAYER, 0)
	assert.Equal(t, rules.SECOND_PLAYER, GetCurrentMark(board))
}

func TestCurrentPlayerIsFirstPlayerAfterTwoMoves(t *testing.T) {
	board := board.NewBoard(rules.THREE_BY_THREE)
	testhelper.AddMarkToPositions(board, rules.FIRST_PLAYER, 0)
	testhelper.AddMarkToPositions(board, rules.SECOND_PLAYER, 1)
	assert.Equal(t, rules.FIRST_PLAYER, GetCurrentMark(board))
}

func TestCurrentPlayerIsFirstPlayerForFirstMoveOnFourByFour(t *testing.T) {
	board := board.NewBoard(rules.FOUR_BY_FOUR)
	testhelper.AddMarkToPositions(board, rules.FIRST_PLAYER, 0)
	assert.Equal(t, rules.SECOND_PLAYER, GetCurrentMark(board))
}

func TestCurrentPlayerIsSecondPlayerOnFourByFour(t *testing.T) {
	board := board.NewBoard(rules.FOUR_BY_FOUR)
	assert.Equal(t, rules.FIRST_PLAYER, GetCurrentMark(board))
}

func TestGetOpponentMark(t *testing.T) {
	assert.Equal(t, rules.SECOND_PLAYER, GetOpponentMark(rules.FIRST_PLAYER))
}

func TestGetOpponentMarkSecondPlayer(t *testing.T) {
	assert.Equal(t, rules.FIRST_PLAYER, GetOpponentMark(rules.SECOND_PLAYER))
}
