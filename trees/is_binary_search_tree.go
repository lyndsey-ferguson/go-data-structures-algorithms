package trees

import "cmp"

/*
Implement a function to check if a binary tree is a binary search tree.

A tree with one node is a binary tree.
A tree with no nodes is a binary tree.

A tree that has a left node that has a value less than or equal to the root may be a binary tree.
If it has a left node that is greater, then it is not.

A tree that has a right node that has a value greater than the root may be a binary tree.
If it has a right node that is less than or equal to the root, then it is not.
*/
func IsBinarySearchTree[T cmp.Ordered](root *Node[T]) bool {
	if root == nil {
		return true
	}
	// let's look at the children before exploring deeper so we can exit early if it is obvious that this
	// is not a binary search tree
	if (root.left != nil && root.left.data > root.data) || (root.right != nil && root.right.data <= root.data) {
		return false
	}
	isBstLeft := IsBinarySearchTree(root.left)
	if !isBstLeft {
		// return early if it is obvious that the left subtree is not a binary search tree
		return false
	}
	isBstRight := IsBinarySearchTree(root.right)
	return isBstRight
}
