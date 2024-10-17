package hashmaps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxNumberOfBalloons(t *testing.T) {
	// "balloonballoon" -> 2
	// "bbaall" -> 0
	// "balloonballoooon" -> 2
	// "balloonnolab" -> 1

	type BalloonSample struct {
		word          string
		expectedCount int
	}
	testCases := []BalloonSample{
		{"balloonballoon", 2},
		{"bbaall", 0},
		{"balloonballoooon", 2},
		{"balloonnolab", 1},
	}
	for _, tc := range testCases {
		assert.Equal(t, tc.expectedCount, maxNumberOfBalloons(tc.word))
	}
}
