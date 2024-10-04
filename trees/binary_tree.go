package trees

import "cmp"

func InsertIntoBinaryTree[T cmp.Ordered](root *Node[T], v T) {
	newNode := CreateNode(v)

	current := root
	var parent *Node[T]

	for current != nil {
		parent = current
		if v < current.data {
			current = current.left
		} else {
			current = current.right
		}
	}

	if v < parent.data {
		parent.left = newNode
	} else {
		parent.right = newNode
	}
}

func CreateBinaryTree[T cmp.Ordered](values []T) *Node[T] {
	var root *Node[T]

	for _, v := range values {
		if root == nil {
			root = CreateNode(v)
		} else {
			InsertIntoBinaryTree(root, v)
		}
	}

	return root
}
