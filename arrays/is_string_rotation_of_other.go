package arrays

import "strings"

/*
Assume you have a method isSubstring which checks if one word is a substring of another.
Given two strings, s1 and s2, write code to check if s2 is a rotation of s1 using only
one call to isSubstring(e.g. "atterbottle" is a rotation of "erbottlwat").

What is the rule that determines if a string is a rotation of anotther?

We could sort both s1 & s2 and compare with is substring, but that feels wrong.

So, the thing that is odd, that calls my attention, is the fact that we're allowed
to use ONE call to isSubstring to solve this.

These two strings are the same length, and have the same letters, just in different
order. Ah, the order difference is not random.

The same word exists in both words, just in a different order. Almost as if it is in
a doubly linked list:

waterbottle
^         |
|----------
    erbottlewat


battlebottle
bottlebattle

battlebottlebattlebottle
      bottlebattle

Can I just create a bigger string with the first string doubled and then
check if s2 is a substring of s1?

watterbottlewaterbottle
    erbottlewat

That checks out.
*/

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
