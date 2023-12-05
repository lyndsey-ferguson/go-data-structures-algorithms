package arrays

/*
There are three types of edits that can be performed on strings:
Insert a character, remove a character, or replace a character.
Given two strings, write a function to check if they are one edit (or zero edits) away.

for exmaple:
  - pale, ple √ 1 char removed
  - pales, pale √ 1 char removed
  - pale, bale √ 1 char replaced
  - pale, bake X 2 chars replaced
  - pae, pale √ 1 char removed

This has an O(n) runtime
*/
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
				// we are in the state where we ALREADY found 1 letter that doesn't
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
