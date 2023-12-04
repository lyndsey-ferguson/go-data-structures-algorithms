package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringHasOnlyUniqueLettersUsingHash(t *testing.T) {
	assert.True(t, stringHasOnlyUniqueLettersUsingHash([]byte("abcdefg")))
	assert.False(t, stringHasOnlyUniqueLettersUsingHash([]byte("abbdefg")))
}

func TestStringHasOnlyUniqueLettersUsingQuicksort(t *testing.T) {
	assert.True(t, stringHasOnlyUniqueLettersUsingQuicksort([]byte("abcdefg")))
	assert.False(t, stringHasOnlyUniqueLettersUsingQuicksort([]byte("abbdefg")))
}

func TestStringHasOnlyUniqueLettersUsingBitVector(t *testing.T) {
	assert.True(t, stringHasOnlyUniqueLettersUsingBitVector([]byte("abcdefg")))
	assert.False(t, stringHasOnlyUniqueLettersUsingBitVector([]byte("abbdefg")))
}
