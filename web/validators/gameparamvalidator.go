package validators

import (
	"errors"
	"strconv"
	"strings"

	"github.com/esanmiguelc/gottt/core/rules"
)

func ValidateParams(
	playerOneType,
	playerTwoType,
	boardSize,
	movesPlayed string,
) (string, string, int, []int, error) {
	boardSizeAsInt, boardError := strconv.Atoi(boardSize)

	if isValidPlayerType(playerOneType) || isValidPlayerType(playerTwoType) {
		return errorResult()
	}

	if boardError != nil || boardSizeAsInt != rules.THREE_BY_THREE {
		return errorResult()
	}

	if strings.ContainsAny(movesPlayed, "9") {
		return errorResult()
	}
	movesPlayedAsInt := convertMovesPlayed(movesPlayed)
	movesPlayedAsInt = removeDuplicates(movesPlayedAsInt)
	return playerOneType, playerTwoType, boardSizeAsInt, movesPlayedAsInt, nil
}

func errorResult() (string, string, int, []int, error) {
	return "", "", 0, []int{}, errors.New("Oops!, something went wrong!")
}

func isValidPlayerType(playerType string) bool {
	return playerType != rules.HUMAN && playerType != rules.COMPUTER
}

func convertMovesPlayed(moves string) []int {
	splitMoves := strings.Split(moves, "")
	var movesPlayed []int
	for _, elem := range splitMoves {
		element, _ := strconv.Atoi(elem)
		movesPlayed = append(movesPlayed, element)
	}

	return movesPlayed
}

func removeDuplicates(moves []int) []int {
	result := []int{}
	seen := map[int]int{}
	for _, value := range moves {
		if _, ok := seen[value]; !ok {
			result = append(result, value)
			seen[value] = value
		}
	}
	return result
}
