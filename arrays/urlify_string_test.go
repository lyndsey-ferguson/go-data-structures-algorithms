package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUrlify(t *testing.T) {
	s1 := []byte("Mr John Smith    ")
	urlify(s1, 13)
	assert.Equal(t, []byte("Mr%20John%20Smith"), s1)

	s2 := []byte("Blinkey")
	urlify(s2, len(s2))
	assert.Equal(t, []byte("Blinkey"), s2)
}
