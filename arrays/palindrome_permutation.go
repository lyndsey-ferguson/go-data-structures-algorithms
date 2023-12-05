package arrays

import "strings"

/*
Given a string, write a function to check if it is a permutation of a palindrome.
A palindrome is a word or phrase that is the same forwards and backwards.
A permutation is a rearrangement of letters.
The palindrome does not need to be limited to just dictionary words.


A palindrome is a word or phrase that is the same forwards or backwards.

The given example is 'Taco cat' because it could be 'Taco cat" or 'tac ocat'

Another palindrome is "Do geese see God?". We do not pay attention to case, spaces, or punctuation.

In this example, I made the observation:

- there are 2 Ds
- there are 2 os
- there are 2 gs
- there are 5 es

In the previous example "Taco cat", we had an even count of all characters except for
1 character: 1 o, 2 ts, 2 cs, and 2as.

o by itself is a palindrome. It is the same forwards and backwards.

oi is not a palindrome, nor a permutation of one.

oio is a palindrome and one permutation.

ooi is a permutation of the previous palindrome.

So, questions:
1. does it work with an even count of duplicate chars?
2. Further, does it work with only 1 set of letters that have an odd count?

Can we iterate amoung letters counting how many of each character, excluding punctuation?

O(n) with space complexity of O(n) if each character of the string is unique

*/

func isAlphaNumeric(b byte) bool {
	if (b < 'a' || b > 'z') && (b < 'A' || b > 'Z') && (b < '0' || b > '9') {
		return false
	}
	return true
}

func isPermutationOfPalindrome(s []byte) bool {
	s = []byte(strings.ToLower(string(s)))
	countOfOdds := 0
	letterCountsHash := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		letter := s[i]
		if !isAlphaNumeric(letter) {
			// we can skip non-alphanumeric characters
			continue
		}

		letterCountsHash[letter] += 1
		if letterCountsHash[letter]%2 == 0 {
			countOfOdds -= 1
		} else {
			countOfOdds += 1
		}
	}

	return (countOfOdds <= 1)
}

/*
What if we cannot use an extra data structure? No hashmap?

My first thought is that we can sort the string, then, with a sliding window
of two, delete each pair and count the number of characters deleted.
Then, if the difference between the letters deleted, and the actual length is
less than one, the string is a permutation of a palindrome.

O(n log n) due to the quicksort + O(n)
*/
func isPermutationOfPalindromeNoExtra(s []byte) bool {
	s = []byte(strings.ToLower(string(s)))
	discardedLetterCount := 0
	QuickSort(s, 0, len(s)-1)

	for i := 0; i < len(s); i++ {
		if !isAlphaNumeric(s[i]) {
			discardedLetterCount++
		} else if i < len(s)-1 && s[i] == s[i+1] {
			discardedLetterCount += 2
			i += 1
		}
	}
	return (len(s)-discardedLetterCount <= 1)
}

func getCaseInsensitiveCharCode(b byte) int {
	if b >= 'A' && b <= 'Z' {
		b = (b - 'A') + 'a'
	}
	if b >= 'a' && b <= 'z' {
		return int(b) - int('a')
	}
	return -1
}

/*
However, we may be able to get to the O(n) runtime if we know that
there are less than 32 character values by using a bit vector.

Like the hashmap, go through each character and toggle the bit of each odd.
If 1 bit or fewer is set to 1, this is a permutation of a palindrome.
*/
func isPermutationOfPalindromeBitVector(s []byte) bool {
	bv := 0

	for i := 0; i < len(s); i++ {
		code := getCaseInsensitiveCharCode(s[i])
		if code != -1 {
			bv ^= (1 << code)
		}
	}
	return (bv&(bv-1) == 0)
}
