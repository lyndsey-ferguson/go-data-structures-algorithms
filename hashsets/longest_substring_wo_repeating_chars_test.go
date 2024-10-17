package hashsets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	assert.Equal(t, 6, lengthOfLongestSubstring("abcdaef"))

	assert.Equal(t, 1, lengthOfLongestSubstring("aaaaa"))
	assert.Equal(t, 10, lengthOfLongestSubstring("abrkaabcdefghijjxxx"))
	assert.Equal(t, 3, lengthOfLongestSubstring("pwwkew"))
}
