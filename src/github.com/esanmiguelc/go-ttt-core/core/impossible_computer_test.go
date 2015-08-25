package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestScoreWhenNeitherPlayerWins(t *testing.T) {
	slots := []string{"", "", "", ""}
	assert.Equal(t, 0, Score(slots))
}
