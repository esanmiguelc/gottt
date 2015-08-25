package core

const THREE_BY_THREE = 3
const FIRST_PLAYER = "X"

func CreateBoard(size int) Board {
	return NewBoard(size)
}

func AddMarkToPositions(board Board, mark string, positions ...int) {
	for _, position := range positions {
		board.Slots[position] = mark
	}
}
