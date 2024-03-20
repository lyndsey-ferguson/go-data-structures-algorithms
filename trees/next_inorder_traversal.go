package trees

type ParentedNode[T comparable] struct {
	data   T
	left   *ParentedNode[T]
	right  *ParentedNode[T]
	parent *ParentedNode[T]
}

func CreateParentedNode[T comparable](data T, childNodes ...*ParentedNode[T]) *ParentedNode[T] {
	n := ParentedNode[T]{
		data:   data,
		parent: nil,
	}
	if childrenCount := len(childNodes); childrenCount > 0 {
		n.left = childNodes[0]
		n.left.parent = &n
		if childrenCount > 1 {
			n.right = childNodes[1]
			n.right.parent = &n
		}
	}
	return &n
}

func NextInOrderTraversal[T comparable](node *ParentedNode[T]) *ParentedNode[T] {
	if node == nil {
		return node
	}
	if node.left == nil && node.parent.left == node {
		return node.parent
	}
	if node.right != nil {
		return node.right
	}
	for candidate := node.parent; candidate != nil; candidate = candidate.parent {
		if candidate.parent != nil && candidate.parent.left == candidate {
			return candidate.parent
		}
	}
	return nil
}
