package trees

import "math"

/*
Implement a function to check if a binary tree is balanced.
For the purpose of this question, a balanced tree is defined
to be a tree such that the heights of the two subtrees of any
node never differ by more than one.
*/
func (root *Node[T]) GetBalancedTreeHeight() (int, bool) {
	if root == nil {
		return 0, true
	}

	leftHeight := 0
	rightHeight := 0
	leftBalanced := false
	rightBalanced := false

	leftHeight, leftBalanced = root.left.GetBalancedTreeHeight()
	if leftBalanced {
		rightHeight, rightBalanced = root.right.GetBalancedTreeHeight()
	}
	return 1 + int(math.Max(float64(leftHeight), float64(rightHeight))),
		math.Abs(float64(leftHeight)-float64(rightHeight)) <= 1 && leftBalanced && rightBalanced
}
