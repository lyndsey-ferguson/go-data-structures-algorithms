package linkedlists

/*
Implement a function to check if a linked list is a palindrome.

A palindrome list would look like:

a -> b -> b -> a

OR

a -> b -> c -> b -> a

It is like we want to make two lists, each list half of the given list, and then compare them.

we could put each element in a stack, with the last elements at the top, then
we could re-walk the list, and pop each element in the stack and compare the nodes.

We can make it even smarter by counting the number of elements put in the stack
and then, as we're re-walking the the list and popping the stack, we can
stop comparing when we've passed (n/2) + 1 nodes.

*/

import "math"

func isListPalindromeUsingStack[T comparable](list *Node[T]) bool {
	var stack Stack[*Node[T]]

	nodeCount := 0
	for cursor := list; cursor != nil; cursor = cursor.next {
		stack.Push(cursor)
		nodeCount++
	}
	midWayIndex := math.Ceil(float64(nodeCount / 2))
	for index, cursor := 0, list; index < int(midWayIndex) && cursor != nil; index, cursor = index+1, cursor.next {
		top, _ := stack.Pop()
		if cursor.data != top.data {
			return false
		}
	}
	return true
}
