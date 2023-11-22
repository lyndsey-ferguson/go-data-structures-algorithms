package linkedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsListPalindromeUsingStackValidWithEvenPalindrome(t *testing.T) {
	list := createLetterList([]string{"a", "b", "b", "a"})

	assert.True(t, isListPalindromeUsingStack(list))
}

func TestIsListPalindromeUsingStackValidWithOddPalindrome(t *testing.T) {
	list := createLetterList([]string{"a", "b", "c", "b", "a"})

	assert.True(t, isListPalindromeUsingStack(list))
}

func TestIsListPalindromeUsingStackInvalid(t *testing.T) {
	list := createLetterList([]string{"a", "b", "b", "c", "a"})

	assert.False(t, isListPalindromeUsingStack(list))
}
