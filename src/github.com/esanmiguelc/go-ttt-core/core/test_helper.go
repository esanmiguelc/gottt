package core

const THREE_BY_THREE = 3
const FIRST_PLAYER = "X"

func CreateSlots(size int) []string {
	return make([]string, size*size)
}

func AddMarkToPositions(slots []string, mark string, positions ...int) {
	for _, position := range positions {
		slots[position] = mark
	}
}
