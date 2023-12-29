package linkedlists

import "cmp"

/*
Write a program to sort a stack such that the smallest items are on the top. You can
use an additional temporary stack, but you may not copy the elements into any other
data structure (such as an array). The stack supports the following operations: push,
pop, peek, and isEmpty.

I could simply push the element if the top of the stack is ermpty or equal or greater
than the pushed value. If the top of the stack is smaller, I pop from that stack and
I push to the temp stack until the top of the stack is empty or greater than or equal to
the pushed value, and then push the given value, and then pop from the temp stack to the
stack until the temp stack is empty.

Otherwise, pop, peek, and isEmpty are all the same as a normal stack.
*/
type SortedStack[T cmp.Ordered] struct {
	Stack[T]
}

func (s *SortedStack[T]) Push(v T) {
	var t Stack[T] // the temporary stack

	e, ok := s.Peek()
	for ; ok && e < v; e, ok = s.Peek() {
		e, ok = s.Pop()
		if ok {
			t.Push(e)
		}
	}
	s.Stack.Push(v)
	for e, ok = t.Pop(); ok; e, ok = t.Pop() {
		s.Stack.Push(e)
	}
}
