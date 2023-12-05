package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCompressedStringReturnsOriginalWhenCompressedStringSameLength(t *testing.T) {
	s := []byte("aabbccdd")
	assert.Equal(t, s, getCompressedString(s))
}

func TestGetCompressedStringReturnsCompressedStringWhenSmaller(t *testing.T) {
	s := []byte("aabcccccaaa")
	assert.Equal(t, []byte("a2bc5a3"), getCompressedString(s))
}
