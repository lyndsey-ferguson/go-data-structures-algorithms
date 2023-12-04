package arrays

/*
Implement an algorithm that can determine if a string consists of all unique characters.
That is to say, each character in the string is the only instance of that character.

For this exercise, we have to iterate at least once through the string, so the fastest runtime is O(n).

Probably, the fastest way we could get this done is by creating a hashmap of each character visited.
If a character is found that already exists, we exit the function with false.

This has a runtime of O(n) and a space complexity of O(n)
*/

func stringHasOnlyUniqueLettersUsingHash(s []byte) bool {
	if len(s) == 0 {
		return false
	}
	if len(s) == 1 {
		return true
	}

	foundLetters := make(map[byte]bool, len(s))
	for index := 0; index < len(s); index++ {
		_, ok := foundLetters[s[index]]
		if ok {
			return false
		} else {
			foundLetters[s[index]] = true
		}
	}
	return true
}

/*
What if I cannot use an additional data structure?
That is to say, what if I'm not allowed a hashmap?

I could presort the string and look at each character's neighbor for a duplicate.

We cannot use mergesort as that would be disallowed by the requirement that we not use an additional data structure:
mergesort uses extra arrays for each subarray.
So that would mean we should use quicksort to sort the array in place.
Both are similar runtimes, O(n log n), but quicksort has a worse worst case scenario if the array is already sorted.

The result of this new method, not using extra data structures is O(n log n)

The quicksort is O(n log n) and the iteration of the sorted array is O(n).
O(n) + O(n log n) is basically O(n log n).
*/

func stringHasOnlyUniqueLettersUsingQuicksort(s []byte) bool {
	if len(s) == 0 {
		return false
	}
	if len(s) == 1 {
		return true
	}
	QuickSort(s, 0, len(s)-1)
	for index := 0; index < len(s)-1; index++ {
		if s[index] == s[index+1] {
			return false
		}
	}
	return true
}

/*
There is an intermediate step, where we are not really using an extra data structure.
Instead, if we can assume that there are only 26 letters, we can use a 32 bit number.
Each bit in the vector can be used to store the code value for each character instead
of a hash and we still get O(n) runtime. That may, or may not be acceptable.
*/

func hashKeyForLetter(c byte) int32 {
	shiftLeftValue := 31
	if c >= byte('A') && c <= byte('Z') {
		shiftLeftValue = int(c - byte('A'))
	} else if c >= byte('a') && c <= byte('z') {
		shiftLeftValue = int(c - byte('a'))
	}
	return (1 << shiftLeftValue)
}

func stringHasOnlyUniqueLettersUsingBitVector(s []byte) bool {
	if len(s) == 0 {
		return false
	}
	if len(s) == 1 {
		return true
	}

	var foundLetters int32
	for index := 0; index < len(s); index++ {
		hashKey := hashKeyForLetter(s[index])
		if foundLetters&hashKey > 0 {
			return false
		} else {
			foundLetters |= hashKey
		}
	}
	return true
}
