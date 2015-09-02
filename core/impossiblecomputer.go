package core

var move int

const maxInt = int(^uint(0) >> 1)
const minInt = -maxInt - 1

type node struct {
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
	minimax(board, myMark, opponent, minInt, maxInt)
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
	if IsGameOver(board) {
		return Score(board, myMark, opponent)
	}
	possibleMoves := []node{}

	for _, position := range MovesAvailable(board) {
		board.PlaceMove(position, GetCurrentMark(board))
		score := minimax(board, myMark, opponent, minValue, maxValue)
		currentNode := node{Score: score, Position: position}

		board.Undo(position)

		if IsCurrentPlayer(board, myMark) {
			possibleMoves = append(possibleMoves, currentNode)
			if currentNode.Score > minValue {
				minValue = currentNode.Score
			}
		} else {
			if currentNode.Score < maxValue {
				maxValue = currentNode.Score
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

func bestNode(nodes []node) node {
	var result node
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
