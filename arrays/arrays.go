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
