package players_test

import (
	"testing"

	"github.com/esanmiguelc/gottt/core/board"
	. "github.com/esanmiguelc/gottt/core/players"
	"github.com/esanmiguelc/gottt/core/rules"
	"github.com/esanmiguelc/gottt/core/testhelper"
	"github.com/stretchr/testify/assert"
)

func TestScoreWhenNeitherPlayerWins(t *testing.T) {
	board := board.NewBoard(rules.THREE_BY_THREE)
	assert.Equal(t, 0, Score(board, rules.FIRST_PLAYER, rules.SECOND_PLAYER))
}

func TestScoreWhenIWin(t *testing.T) {
	board := board.NewBoard(rules.THREE_BY_THREE)
	testhelper.AddMarkToPositions(board, rules.FIRST_PLAYER, 0, 1, 2)
	assert.Equal(t, 10, Score(board, rules.FIRST_PLAYER, rules.SECOND_PLAYER))
}

func TestScoreWhenOpponentWins(t *testing.T) {
	board := board.NewBoard(rules.THREE_BY_THREE)
	testhelper.AddMarkToPositions(board, rules.SECOND_PLAYER, 0, 1, 2)
	assert.Equal(t, -10, Score(board, rules.FIRST_PLAYER, rules.SECOND_PLAYER))
}

func TestGetsTheFirstMove(t *testing.T) {
	board := board.NewBoard(rules.THREE_BY_THREE)
	impossiblePlayer := ImpossiblePlayer{Mark: rules.SECOND_PLAYER}
	assert.Equal(t, 0, impossiblePlayer.GetMove(board, rules.FIRST_PLAYER, rules.SECOND_PLAYER))
}

func TestGetsTheCenterMove(t *testing.T) {
	board := board.NewBoard(rules.THREE_BY_THREE)
	impossiblePlayer := ImpossiblePlayer{Mark: rules.SECOND_PLAYER}
	testhelper.AddMarkToPositions(board, rules.FIRST_PLAYER, 0)
	assert.Equal(t, 4, impossiblePlayer.GetMove(board, rules.SECOND_PLAYER, rules.FIRST_PLAYER))
}

func TestBlocksAWin(t *testing.T) {
	board := board.NewBoard(rules.THREE_BY_THREE)
	impossiblePlayer := ImpossiblePlayer{Mark: rules.SECOND_PLAYER}
	testhelper.AddMarkToPositions(board, rules.FIRST_PLAYER, 0, 2, 5, 7)
	testhelper.AddMarkToPositions(board, rules.SECOND_PLAYER, 1, 4, 3)
	assert.Equal(t, 8, impossiblePlayer.GetMove(board, rules.SECOND_PLAYER, rules.FIRST_PLAYER))
}

func TestGoesForTheWin(t *testing.T) {
	board := board.NewBoard(rules.THREE_BY_THREE)
	impossiblePlayer := ImpossiblePlayer{Mark: rules.SECOND_PLAYER}
	testhelper.AddMarkToPositions(board, rules.FIRST_PLAYER, 0, 5, 2)
	testhelper.AddMarkToPositions(board, rules.SECOND_PLAYER, 1, 4)
	assert.Equal(t, 7, impossiblePlayer.GetMove(board, rules.SECOND_PLAYER, rules.FIRST_PLAYER))
}

func TestNewImpossibleComputer(t *testing.T) {
	impossiblePlayer := ImpossiblePlayer{Mark: rules.FIRST_PLAYER}
	assert.True(t, impossiblePlayer.IsComputer())
}
