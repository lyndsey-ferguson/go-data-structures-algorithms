package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArePermutationsFailsWithUnequalLengthStrings(t *testing.T) {
	result := ArePermutations("abcde", "abcdefg")
	assert.False(t, result)
}

func TestArePermutationsPassesReorderedEqualStrings(t *testing.T) {
	result := ArePermutations("aaabcca", "aaaacbc")
	assert.True(t, result)
}

func TestArePermutationsFailsWithDifferentStrings(t *testing.T) {
	result := ArePermutations("aaabcca", "aaaabbc")
	assert.False(t, result)
}

func TestUrlify(t *testing.T) {
	s1 := []byte("Mr John Smith    ")
	urlify(s1, 13)
	assert.Equal(t, []byte("Mr%20John%20Smith"), s1)

	s2 := []byte("Blinkey")
	urlify(s2, len(s2))
	assert.Equal(t, []byte("Blinkey"), s2)
}

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
