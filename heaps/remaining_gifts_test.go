package heaps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemainingGifts(t *testing.T) {
	assert.Equal(t, 11, remainingGifts([]int{4, 9, 16}, 2))
}
