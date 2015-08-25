package tttboard

type Board struct {
	Slots []string
}

func createBoard(size int) Board {
	return Board{make([]string, size*size)}
}

func (board Board) PlaceMove(position int, mark string) {
	board.Slots[position] = mark
}
