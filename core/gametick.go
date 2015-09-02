package core

func GameTick(playerOneType, playerTwoType string, boardSize int, movesPlayed []int) GameState {
	playerOne := createFirstPlayer(playerOneType)
	playerTwo := createSecondPlayer(playerTwoType)
	board := NewBoard(boardSize)
	rebuildBoard(board, movesPlayed)

	return GameState{
		MovesPlayed:   takeTurn(board, playerOne, playerTwo, movesPlayed),
		PlayerOneType: playerOneType,
		PlayerTwoType: playerTwoType,
		CurrentPlayer: GetCurrentPlayer(board, playerOne, playerTwo),
		BoardSize:     boardSize,
		BoardState:    board.Slots,
		Board:         board,
	}
}

func rebuildBoard(board Board, movesPlayed []int) {
	for _, move := range movesPlayed {
		board.PlaceMove(move, GetCurrentMark(board))
	}
}

func takeTurn(board Board, playerOne, playerTwo Player, movesPlayed []int) []int {
	currentPlayer := GetCurrentPlayer(board, playerOne, playerTwo)
	if currentPlayer.IsComputer() {
		move := currentPlayer.GetMove(board, currentPlayer.GetMark(), GetOpponentMark(currentPlayer.GetMark()))
		board.PlaceMove(move, currentPlayer.GetMark())
		movesPlayed = append(movesPlayed, move)
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
	if playerType == COMPUTER {
		return ImpossiblePlayer{Mark: mark}
	}
	return HumanPlayer{mark}
}
