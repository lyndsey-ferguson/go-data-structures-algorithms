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
