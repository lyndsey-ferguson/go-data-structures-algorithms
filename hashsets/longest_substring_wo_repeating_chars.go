package hashsets

import "fmt"

type Solution struct{}

func lengthOfLongestSubstring(str string) int {
	/**
	we add unique chars to hashset keeping track of start and end
	as we loop through str we check:
	- is the string at the "end" unique?
		- if so, add it and increase the end and update the longest length val if
			end - start + 1 is bigger
		- if not, remove the char from the str's "start" and increment start as we've
			discovered that up to and including that char, we have a duplicated string
	- at the end, return the found longest length
	*/
	set := make(map[byte]bool)

	longest := 0
	start := 0
	end := 0
	for end < len(str) {
		if set[str[end]] {
			fmt.Printf("str[start] (%v) !=? str[end] (%v)\n", str[start], str[end])
			delete(set, str[start])
			start++
		} else {
			set[str[end]] = true
			if end-start+1 > longest {
				longest = end - start + 1
			}
			end++
		}
	}
	return longest
}
