package arrays

import "strconv"

/*
Implement a method to perform basic string compression using the counts of repeated characters.
For example, the string aabcccccaaa would become a2b1c5a3.
If the "compressed" string would not become smaller than the original string,
your method should return the original string.
You can assume the string has only uppercase and lowercase letters (a-z).

We can scan the entire string to see if there is a dup. If there are none
return the original.
*/
func insertNumber(s []byte, number int, index *int) {
	numStr := strconv.Itoa(number)
	for k := 0; k < len(numStr); k++ {
		s[*index] = numStr[k]
		(*index)++
	}
}

func getCompressedString(s []byte) []byte {
	if len(s) < 3 {
		// there are only 2 characters, the best we could do would be if
		// those 2 characters are the same, but then resultant 'compressed'
		// string is still 2 characters, for example, aa -> a1 same len.
		return s
	}
	var lastChar byte
	charCount := 0
	compressed := make([]byte, len(s))
	j := 0
	for i := 0; i < len(s); i++ {
		// if the character is the same as the last one, we can track
		// the number of characters, otherwise, we reset the character
		// count
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
