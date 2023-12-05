package arrays

import (
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
