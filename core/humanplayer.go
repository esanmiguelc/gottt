package core

type HumanPlayer struct {
	Mark string
}

func (human HumanPlayer) GetMove(board Board, myMark, opponent string) int {
	return 0
}

func (human HumanPlayer) GetMark() string {
	return human.Mark
}

func (human HumanPlayer) IsComputer() bool {
	return false
}

func NewHumanPlayer(mark string) HumanPlayer {
	return HumanPlayer{Mark: mark}
}
