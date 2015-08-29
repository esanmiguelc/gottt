package core

type Player interface {
	GetMark() string
	IsComputer() bool
	GetMove(board Board, myMark, opponent string) int
}
