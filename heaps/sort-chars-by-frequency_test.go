package heaps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFrequencySort(t *testing.T) {
	assert.Equal(t, "aaannb", frequencySort("banana"))
}
