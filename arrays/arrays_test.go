package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeSort(t *testing.T) {
	array := []int{12, 11, 13, 5, 6, 7}
	mergedArray := array

	MergeSort(&mergedArray, 0)
	assert.Equal(t, []int{5, 6, 7, 11, 12, 13}, mergedArray)
}

func TestQuickSort(t *testing.T) {
	array := []int{10, 7, 8, 9, 1, 5}
	QuickSort(array, 0, len(array)-1)
	assert.Equal(t, []int{1, 5, 7, 8, 9, 10}, array)
}

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
