package trees

import "math"

/*
Given a sorted (increasing order) array with unique integer elements
write an algorithm to create a binary search tree with minimal height.
*/

func TreeifySortedArray[T comparable](a []T, start int, end int) *Node[T] {
	var root *Node[T]

	if end < start {
		return nil
	} else if (end - start) <= 2 {
		root = CreateNode(a[start])
		if end > start {
			root.right = CreateNode(a[end])
		}
	} else {
		rootIndex := start + int(math.Ceil(float64(end-start)/2))
		root = CreateNode(a[rootIndex])
		root.left = TreeifySortedArray(a, start, rootIndex-1)
		root.right = TreeifySortedArray(a, rootIndex+1, end)
	}
	return root
}
