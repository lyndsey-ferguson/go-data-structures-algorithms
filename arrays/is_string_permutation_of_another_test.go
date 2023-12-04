package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArePermutationsFailsWithUnequalLengthStrings(t *testing.T) {
	result := ArePermutations("abcde", "abcdefg")
	assert.False(t, result)
}

func TestArePermutationsPassesReorderedEqualStrings(t *testing.T) {
	result := ArePermutations("aaabcca", "aaaacbc")
	assert.True(t, result)
}

func TestArePermutationsFailsWithDifferentStrings(t *testing.T) {
	result := ArePermutations("aaabcca", "aaaabbc")
	assert.False(t, result)
}
