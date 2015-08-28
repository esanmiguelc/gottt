package core

type GameState struct {
	MovesPlayed   []int
	PlayerOneType string
	PlayerTwoType string
	BoardSize     int
}

func GameTick(playerOneType, playerTwoType string, boardSize int, movesPlayed []int) GameState {
	playerOne := createFirstPlayer(playerOneType)
	playerTwo := createSecondPlayer(playerTwoType)
	board := NewBoard(boardSize)

	rebuildBoard(board, movesPlayed)

	return GameState{
		MovesPlayed:   takeTurn(board, playerOne, playerTwo, movesPlayed),
		PlayerOneType: playerOneType,
		PlayerTwoType: playerTwoType,
		BoardSize:     boardSize,
	}
}

func rebuildBoard(board Board, movesPlayed []int) {
	for _, move := range movesPlayed {
		board.PlaceMove(move, GetCurrentPlayer(board))
	}
}

func takeTurn(board Board, playerOne, playerTwo Player, movesPlayed []int) []int {
	if IsCurrentPlayer(board, playerOne.GetMark()) {
		if playerOne.IsComputer() {
			move := playerOne.GetMove(board, playerOne.GetMark(), playerTwo.GetMark())
			movesPlayed = append(movesPlayed, move)
		}
	} else {
		if playerTwo.IsComputer() {
			move := playerTwo.GetMove(board, playerTwo.GetMark(), playerOne.GetMark())
			movesPlayed = append(movesPlayed, move)
		}
	}
	return movesPlayed
}

func createFirstPlayer(playerType string) Player {
	return createPlayer(playerType, FIRST_PLAYER)
}

func createSecondPlayer(playerType string) Player {
	return createPlayer(playerType, SECOND_PLAYER)
}

func createPlayer(playerType, mark string) Player {
	if playerType == "Computer" {
		return ImpossiblePlayer{Mark: mark}
	} else if playerType == "Human" {
		return HumanPlayer{mark}
	} else {
		return HumanPlayer{mark}
	}
}