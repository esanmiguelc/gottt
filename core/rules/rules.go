package rules

import (
	math "math"

	"github.com/esanmiguelc/gottt/core/board"
)

const THREE_BY_THREE = 3
const FOUR_BY_FOUR = 4
const FIRST_PLAYER = "X"
const SECOND_PLAYER = "O"
const TIE = "Tie"
const HUMAN = "Human"
const COMPUTER = "Computer"

func GetResult(board board.Board) string {
	if IsGameOver(board) {
		if IsWinner(board, FIRST_PLAYER) {
			return FIRST_PLAYER
		} else if IsWinner(board, SECOND_PLAYER) {
			return SECOND_PLAYER
		} else {
			return TIE
		}
	}
	return ""
}

func MovesAvailable(board board.Board) []int {
	result := []int{}
	for index, value := range board.Slots {
		if value == "" {
			result = append(result, index)
		}
	}
	return result
}

func IsWinner(board board.Board, mark string) bool {
	return ColumnWinner(board, mark) || DiagonalWinner(board, mark) || RowWinner(board, mark)
}

func IsGameOver(board board.Board) bool {
	return board.Full() || IsWinner(board, FIRST_PLAYER) || IsWinner(board, SECOND_PLAYER)
}

func ColumnWinner(board board.Board, mark string) bool {
	columnCombinations := getPossibleColumnCombinations(board.Slots)
	return checkCombinations(columnCombinations, board.Slots, mark)
}

func DiagonalWinner(board board.Board, mark string) bool {
	diagonalCombinations := getPossibleDiagonalCombinations(board.Slots)
	return checkCombinations(diagonalCombinations, board.Slots, mark)
}

func RowWinner(board board.Board, mark string) bool {
	rowCombinations := getPossibleRowCombinations(board.Slots)
	return checkCombinations(rowCombinations, board.Slots, mark)
}

func allTrue(list []bool) bool {
	for _, value := range list {
		if value != true {
			return false
		}
	}
	return true
}

func checkCombinations(combinations [][]int, slots []string, mark string) bool {
	for _, rowValue := range combinations {
		result := make([]bool, getBoardSize(slots))
		for index, value := range rowValue {
			result[index] = slots[value] == mark
		}
		if allTrue(result) {
			return true
		}
	}
	return false
}

func getBoardSize(slots []string) int {
	return GetBoardSize(slots)
}

func GetBoardSize(slots []string) int {
	return int(math.Sqrt(float64(len(slots))))
}

func getPossibleColumnCombinations(slots []string) [][]int {
	boardSize := getBoardSize(slots)
	possibleColumnCombinations := make([][]int, boardSize)
	for column := 0; column < boardSize; column++ {
		possibleColumnCombinations[column] = make([]int, boardSize)
		for position := 0; position < boardSize; position++ {
			possibleColumnCombinations[column][position] = column + (position * boardSize)
		}
	}
	return possibleColumnCombinations
}

func getPossibleDiagonalCombinations(slots []string) [][]int {
	boardSize := getBoardSize(slots)
	possibleDiagonalCombinations := make([][]int, 2)
	for index := range possibleDiagonalCombinations {
		possibleDiagonalCombinations[index] = make([]int, boardSize)
	}
	for position := 0; position < boardSize; position++ {
		possibleDiagonalCombinations[0][position] = position * (boardSize + 1)
		possibleDiagonalCombinations[1][position] = (position * (boardSize - 1)) + boardSize - 1
	}
	return possibleDiagonalCombinations
}

func getPossibleRowCombinations(slots []string) [][]int {
	boardSize := getBoardSize(slots)
	possibleRowCombinations := make([][]int, boardSize)
	slotPosition := 0
	for row := 0; row < boardSize; row++ {
		possibleRowCombinations[row] = make([]int, boardSize)
		for index := range possibleRowCombinations[row] {
			possibleRowCombinations[row][index] = slotPosition
			slotPosition += 1
		}
	}
	return possibleRowCombinations
}
