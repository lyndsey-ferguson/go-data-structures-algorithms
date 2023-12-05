package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAreStringsLessThanTwoEditsApart(t *testing.T) {
	s1, s2 := []byte("pale"), []byte("ple")
	assert.True(t, areStringsLessThanTwoEditsApart(s1, s2))

	s1, s2 = []byte("pales"), []byte("pale")
	assert.True(t, areStringsLessThanTwoEditsApart(s1, s2))

	s1, s2 = []byte("pale"), []byte("bale")
	assert.True(t, areStringsLessThanTwoEditsApart(s1, s2))

	s1, s2 = []byte("pale"), []byte("bake")
	assert.False(t, areStringsLessThanTwoEditsApart(s1, s2))

	s1, s2 = []byte("pae"), []byte("pale")
	assert.True(t, areStringsLessThanTwoEditsApart(s1, s2))
}

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
