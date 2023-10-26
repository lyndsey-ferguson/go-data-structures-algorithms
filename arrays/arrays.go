package arrays

import (
	"cmp"
	"strconv"
	"strings"
)

func ArrayToString(array []int) string {
	stringArray := []string{}
	for i := range array {
		number := array[i]
		text := strconv.Itoa(number)
		stringArray = append(stringArray, text)
	}
	return "[" + strings.Join(stringArray, " ") + "]"
}

func MergeSort(array *[]int, indent int) {
	arr := *array
	if len(arr) > 1 {
		mid := len(arr) / 2

		L := make([]int, mid)
		copy(L[:], arr[:mid])
		R := make([]int, len(arr)-mid)
		copy(R[:], arr[mid:])
		MergeSort(&L, indent+1)
		MergeSort(&R, indent+1)

		i, j, k := 0, 0, 0
		for i < len(L) && j < len(R) {
			if L[i] <= R[j] {
				arr[k] = L[i]
				i++
			} else {
				arr[k] = R[j]
				j++
			}
			k++
		}
		for i < len(L) {
			arr[k] = L[i]
			i++
			k++
		}
		for j < len(R) {
			arr[k] = R[j]
			j++
			k++
		}
	}
}

func partition[T cmp.Ordered](arr []T, low int, high int) int {
	pivot := arr[high]

	bestKnownCorrectPivotIndex := low - 1
	for index := low; index <= high-1; index++ {
		if arr[index] < pivot {
			bestKnownCorrectPivotIndex++
			arr[bestKnownCorrectPivotIndex], arr[index] = arr[index], arr[bestKnownCorrectPivotIndex]
		}
	}
	arr[bestKnownCorrectPivotIndex+1], arr[high] = arr[high], arr[bestKnownCorrectPivotIndex+1]

	return bestKnownCorrectPivotIndex + 1
}

func QuickSort[T cmp.Ordered](arr []T, low int, high int) {
	if low < high {
		partitionIndex := partition(arr, low, high)
		QuickSort(arr, low, partitionIndex-1)
		QuickSort(arr, partitionIndex+1, high)
	}
}

func ArePermutations(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	lettersCounter := make(map[byte]int)

	for i := 0; i < len(s1); i++ {
		lettersCounter[s1[i]]++
	}
	for i := 0; i < len(s2); i++ {
		lettersCounter[s2[i]]--
		if lettersCounter[s2[i]] < 0 {
			return false
		}
	}
	return true
}

func urlify(s []byte, trueLength int) {
	j := len(s) - 1
	i := trueLength - 1
	for ; i >= 0 && i < j; i-- {
		if s[i] == ' ' {
			s[j] = '0'
			s[j-1] = '2'
			s[j-2] = '%'
			j -= 3
		} else {
			s[j], j = s[i], j-1
		}
	}
}

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

func Abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func areStringsLessThanTwoEditsApart(s1 []byte, s2 []byte) bool {
	if Abs(len(s1)-len(s2)) > 1 {
		return false
	}
	bigger, smaller := s1, s2
	if len(smaller) > len(bigger) {
		bigger, smaller = smaller, bigger
	}
	foundOffChar := false
	j := 0
	for i := 0; i < len(bigger) && j < len(smaller); i++ {
		if bigger[i] == smaller[j] {
			j++
		} else {
			if foundOffChar {
				// we are in the state where we already found 1 letter that doesn't
				// line up. Either the letter is the one that was added to bigger
				// or removed from smaller.
				return false
			} else {
				if len(bigger) == len(smaller) {
					// we may have found the letter that was replaced in one of the
					// strings. Let's update 'j' so that the next comparision
					// will be lined up with the 'ith' character.
					j++
				}
				foundOffChar = true
			}
		}
	}

	return true
}

func insertNumber(s []byte, number int, index *int) {
	numStr := strconv.Itoa(number)
	for k := 0; k < len(numStr); k++ {
		s[*index] = numStr[k]
		(*index)++
	}
}

func getCompressedString(s []byte) []byte {
	if len(s) < 3 {
		return s
	}
	var lastChar byte
	charCount := 0
	compressed := make([]byte, len(s))
	j := 0
	for i := 0; i < len(s); i++ {
		if s[i] != lastChar {
			if charCount > 1 {
				insertNumber(compressed, charCount, &j)
			}
			lastChar = s[i]
			compressed[j] = lastChar
			charCount = 1
			j++
		} else {
			charCount++
		}
	}
	if charCount > 1 {
		insertNumber(compressed, charCount, &j)
	} else {
		compressed[j] = lastChar
	}
	if j < len(s) {
		return compressed[:j]
	}
	return s
}

func isStringRotation(s1 []byte, s2 []byte) bool {
	if len(s1) != len(s2) {
		return false
	}

	s1s1 := make([]byte, len(s1)*2)
	j := 0
	for reps := 0; reps < 2; reps++ {
		for i := 0; i < len(s1); i++ {
			s1s1[j] = s1[i]
			j++
		}
	}
	return strings.Contains(string(s1s1), string(s2))
}
