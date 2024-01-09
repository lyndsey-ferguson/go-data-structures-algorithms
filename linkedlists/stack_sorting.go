package linkedlists

import "cmp"

type ComparableStack[T cmp.Ordered] struct {
	Stack[T]
}

func (left *ComparableStack[T]) Sort() {
	right := ComparableStack[T]{}
	for pivot, ok_l := left.Pop(); ok_l; pivot, ok_l = left.Pop() {
		r, ok_r := right.Peek()
		for ; ok_r && (pivot >= r); r, ok_r = right.Peek() {
			r, _ = right.Pop()
			left.Push(r)
		}
		if pivot < r || !ok_r {
			right.Push(pivot)
		}
	}
	*left = right
}
