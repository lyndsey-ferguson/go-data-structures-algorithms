package trees

import "cmp"

func mergeTrees[T cmp.Ordered](t1 *Node[T], t2 *Node[T]) *Node[T] {
	if t1 == nil {
		return t2
	} else if t2 == nil {
		return t1
	}

	t1.data = t1.data + t2.data

	t1.left = mergeTrees(t1.left, t2.left)
	t1.right = mergeTrees(t1.right, t2.right)

	return t1
}
