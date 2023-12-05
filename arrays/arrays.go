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
