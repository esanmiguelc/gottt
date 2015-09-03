package testhelper

import "github.com/esanmiguelc/gottt/core/board"

func AddMarkToPositions(board board.Board, mark string, positions ...int) {
	for _, position := range positions {
		board.Slots[position] = mark
	}
}
