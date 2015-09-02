package core

func AddMarkToPositions(board Board, mark string, positions ...int) {
	for _, position := range positions {
		board.Slots[position] = mark
	}
}
