package arrays

/*
Given two strings, write a method to decide if one is a permutation of the other string
*/

/*
Let's consider the question, what is a permutation? It means that all the characters in
one string are in the other string. The order is not important.

Here is an example:

- s1 = aaabcca
- s2 = aaaacbc

That is a permutation as all the characters in s1 are in s2.

Here is another example:

- s1 = aabbcc
- s2 = aabbc

This is not a permutation as not all the characters are in s2 that are in s1.

What is a brute force algorithm?

1 - return false if the string lengths are not equal, that's easy. If the lengths
are different, we know that there is no way to have the same number of characters
in one string be in the other.

2 - get a count of each character and make sure that it is the same for both??

OR

2 - use a hash map to get a count of each character and then iterate over the keys
of s1 count and compare to the keys in s2 count. If any don't match, exit by returning
false.

Now look, what are the bottlenecks, unnecessary work, or duplicated work? We are looping through the
two strings, and then looping through the keys in each hashmap.

What if we incremented the string count for a character in s1 in one loop, and decremented the string count for a
character in s2 in another loop, if after any decrement, the count is less than 0, we can exit with false as we
just found that there are more characters of a certain character in s2 than in s1.

That would be two loops O(2n) -> O(n)
*/

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
