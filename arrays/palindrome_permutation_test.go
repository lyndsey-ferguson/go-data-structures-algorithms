package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPermutationOfPalindrome(t *testing.T) {
	s1 := []byte("i")
	assert.True(t, isPermutationOfPalindrome(s1))

	s2 := []byte("oio")
	assert.True(t, isPermutationOfPalindrome(s2))

	s3 := []byte("oi")
	assert.False(t, isPermutationOfPalindrome(s3))

	s4 := []byte("tact coa")
	assert.True(t, isPermutationOfPalindrome(s4))

	s5 := []byte("tact cot")
	assert.False(t, isPermutationOfPalindrome(s5))

	s6 := []byte("A man, a plan, a canal—Panama")
	assert.True(t, isPermutationOfPalindrome(s6))

	s7 := []byte("aa")
	assert.True(t, isPermutationOfPalindrome((s7)))
}

func TestIsPermutationOfPalindromeNoExtra(t *testing.T) {
	s1 := []byte("i")
	assert.True(t, isPermutationOfPalindromeNoExtra(s1))

	s2 := []byte("oio")
	assert.True(t, isPermutationOfPalindromeNoExtra(s2))

	s3 := []byte("oi")
	assert.False(t, isPermutationOfPalindromeNoExtra(s3))

	s4 := []byte("tact coa")
	assert.True(t, isPermutationOfPalindromeNoExtra(s4))

	s5 := []byte("tact cot")
	assert.False(t, isPermutationOfPalindromeNoExtra(s5))

	s6 := []byte("A man, a plan, a canal—Panama")
	assert.True(t, isPermutationOfPalindromeNoExtra(s6))

	s7 := []byte("aa")
	assert.True(t, isPermutationOfPalindromeNoExtra((s7)))
}

func TestIsPermutationOfPalindromeBitVector(t *testing.T) {
	s1 := []byte("i")
	assert.True(t, isPermutationOfPalindromeBitVector(s1))

	s2 := []byte("oio")
	assert.True(t, isPermutationOfPalindromeBitVector(s2))

	s3 := []byte("oi")
	assert.False(t, isPermutationOfPalindromeBitVector(s3))

	s4 := []byte("tact coa")
	assert.True(t, isPermutationOfPalindromeBitVector(s4))

	s5 := []byte("tact cot")
	assert.False(t, isPermutationOfPalindromeBitVector(s5))

	s6 := []byte("A man, a plan, a canal—Panama")
	assert.True(t, isPermutationOfPalindromeBitVector(s6))

	s7 := []byte("aa")
	assert.True(t, isPermutationOfPalindromeBitVector((s7)))
}
