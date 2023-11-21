package linkedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsListPalindromeValidWithEvenPalindrome(t *testing.T) {
	list := createLetterList([]string{"a", "b", "b", "a"})

	assert.True(t, isListPalindrome(list))
}

func TestIsListPalindromeValidWithOddPalindrome(t *testing.T) {
	list := createLetterList([]string{"a", "b", "c", "b", "a"})

	assert.True(t, isListPalindrome(list))
}

func TestIsListPalindromeInvalid(t *testing.T) {
	list := createLetterList([]string{"a", "b", "b", "c", "a"})

	assert.False(t, isListPalindrome(list))
}
