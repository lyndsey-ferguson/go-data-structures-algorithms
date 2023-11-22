package linkedlists

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
