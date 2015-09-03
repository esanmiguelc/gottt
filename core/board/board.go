package board

type Board struct {
	Slots []string
}

func NewBoard(size int) Board {
	return Board{make([]string, size*size)}
}

func (board Board) Full() bool {
	for _, value := range board.Slots {
		if value == "" {
			return false
		}
	}
	return true
}

func (board Board) Undo(position int) {
	board.Slots[position] = ""
}

func (board Board) PlaceMove(position int, mark string) {
	board.Slots[position] = mark
}
