package arrays

/*
Write a method to replace all spaces in a string with '%20'.
You may assume that the string has sufficient space at the end to hold the additional characters,
and that you are given the "true" length of the string.

If using Java, please use a character array so that you can perform this operation in place.

We are not using Java, but Go.

So, what do we know?

 1. We know that the string ends in spaces.
 2. It ends in eough spaces. So, if there is 1 space in the string, the string itself will end with 2 more spaces.
    2 spaces in the string means that it ends in 4 spaces.
 3. We can assume that the challenge is to perform the operation in place by the point about Java.

We set two "pointers": i & j. 'i' points at the actual true end of the string, and 'j' points to the end of the buffer.

We move 'i' back within the string towards the beginning; as each 'i' reads a real character, we move it to 'j' and back
up each pointer. If 'i' hits a space, we fill 'back' 'j' with '%20' and then move both 'i' and 'j' back one character.
*/
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
