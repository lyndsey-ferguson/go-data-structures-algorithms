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

func TestIsStringRotation(t *testing.T) {
	s1 := []byte("waterbottle")
	s2 := []byte("erbottlewat")

	assert.True(t, isStringRotation(s1, s2))

	s3 := []byte("battlebottle")
	s4 := []byte("bottlebattle")

	assert.True(t, isStringRotation(s3, s4))

	s5 := []byte("baterwottle")
	assert.False(t, isStringRotation(s1, s5))
}
