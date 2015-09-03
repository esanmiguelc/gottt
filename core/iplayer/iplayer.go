package iplayer

import "github.com/esanmiguelc/gottt/core/board"

type Player interface {
	GetMark() string
	IsComputer() bool
	GetMove(board board.Board, myMark, opponent string) int
}
