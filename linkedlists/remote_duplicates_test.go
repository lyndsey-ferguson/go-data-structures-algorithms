package linkedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicatesInline(t *testing.T) {
	list := createLetterList([]string{"x", "w", "y", "x", "h", "y", "a", "x"})
	expectedList := createLetterList([]string{"x", "w", "y", "h", "a"})

	removeDuplicatesInline(list)
	assert.Equal(t, expectedList.ToString(), list.ToString())
}

func TestRemoveDuplicatesWithHash(t *testing.T) {
	list := createLetterList([]string{"x", "w", "y", "x", "h", "y", "a", "x"})
	expectedList := createLetterList([]string{"x", "w", "y", "h", "a"})

	removeDuplicatesWithHash(list)
	assert.Equal(t, expectedList.ToString(), list.ToString())
}
