package core

var move int = 0

const MAX_INT = int(^uint(0) >> 1)
const MIN_INT = -MAX_INT - 1

type Node struct {
	Score    int
	Position int
}

type ImpossiblePlayer struct {
	Mark string
}

func (player ImpossiblePlayer) GetMark() string {
	return player.Mark
}

func (player ImpossiblePlayer) GetMove(board Board, myMark, opponent string) int {
	minimax(board, myMark, opponent, MIN_INT, MAX_INT)
	return move
}
func Score(board Board, myMark, opponent string) int {
	if IsWinner(board, myMark) {
		return 10
	} else if IsWinner(board, opponent) {
		return -10
	} else {
		return 0
	}
}

func (player ImpossiblePlayer) IsComputer() bool {
	return true
}

func minimax(board Board, myMark, opponent string, minValue, maxValue int) int {
	if IsGameOver(board, myMark, opponent) {
		return Score(board, myMark, opponent)
	}
	possibleMoves := []Node{}

	for _, position := range MovesAvailable(board) {
		board.PlaceMove(position, GetCurrentPlayer(board))
		score := minimax(board, myMark, opponent, minValue, maxValue)
		node := Node{Score: score, Position: position}

		board.Undo(position)

		if IsCurrentPlayer(board, myMark) {
			possibleMoves = append(possibleMoves, node)
			if node.Score > minValue {
				minValue = node.Score
			}
		} else {
			if node.Score < maxValue {
				maxValue = node.Score
			}
		}
		if maxValue <= minValue {
			break
		}
	}

	if !IsCurrentPlayer(board, myMark) {
		return maxValue
	}

	move = bestNode(possibleMoves).Position
	return minValue
}

func bestNode(nodes []Node) Node {
	var result Node
	for index := range nodes {
		if index == 0 {
			result = nodes[0]
		} else {
			if nodes[index].Score > result.Score {
				result = nodes[index]
			}
		}
	}
	return result
}
