package core

type GameState struct {
	MovesPlayed   []int
	PlayerOneType string
	PlayerTwoType string
	CurrentPlayer Player
	BoardSize     int
	BoardState    []string
	Board         Board
}
