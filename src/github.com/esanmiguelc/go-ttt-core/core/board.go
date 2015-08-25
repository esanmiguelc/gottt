package core

type Board struct {
	Slots []string
}

func NewBoard(size int) Board {
	return Board{make([]string, size*size)}
}

func (board Board) PlaceMove(position int, mark string) {
	board.Slots[position] = mark
}
