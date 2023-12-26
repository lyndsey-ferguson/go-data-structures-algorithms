package linkedlists

import (
	"cmp"
)

type MinStack[T cmp.Ordered] struct {
	top *Node[T]
	min *Node[*Node[T]]
}

func (s *MinStack[T]) Push(i T) {
	n := CreateNode(i)
	m := CreateNode[*Node[T]](n)

	// push node onto stack
	n.next = s.top
	s.top = n

	// iterate in min list and
	// insert the node in its spot
	// in the list
	if s.min == nil || s.min.data.data >= i {
		m.next = s.min
		s.min = m
	} else {
		cursor := s.min
		for ; cursor != nil && cursor.next != nil; cursor = cursor.next {
			if cursor.next.data.data > i {
				m.next = cursor.next
				break
			}
		}
		if cursor != nil {
			cursor.next = m
		}
	}
}

func (s *MinStack[T]) Pop() (T, bool) {
	top := s.top
	if top != nil {
		if s.min.data == top {
			s.min = s.min.next
		} else {
			cursor := s.min
			for ; cursor != nil && cursor.next != nil; cursor = cursor.next {
				if cursor.next.data == top {
					cursor.next = cursor.next.next
				}
			}
		}
		s.top = s.top.next
		top.next = nil
		return top.data, true
	}
	var nonExistant T
	return nonExistant, false
}

func (s *MinStack[T]) Peek() (T, bool) {
	if s.top != nil {
		return s.top.data, true
	}
	var nonExistant T
	return nonExistant, false
}

func (s *MinStack[T]) Min() (T, bool) {
	if s.min != nil {
		return s.min.data.data, true
	}
	var nonExistant T
	return nonExistant, false
}

func (s *MinStack[T]) IsEmpty() bool {
	return s.top == nil
}

type MinStack2[T cmp.Ordered] struct {
	top Stack[T]
	min Stack[T]
}

func (s *MinStack2[T]) Push(i T) {
	s.top.Push(i)
	v, ok := s.min.Peek()
	if !ok || i <= v {
		s.min.Push(i)
	}
}

func (s *MinStack2[T]) Pop() (T, bool) {
	e, ok := s.top.Pop()
	if ok {
		v, ok := s.min.Peek()
		if ok && v == e {
			s.min.Pop()
		}
		return e, ok
	}
	var nonExistant T
	return nonExistant, false
}

func (s *MinStack2[T]) Peek() (T, bool) {
	return s.top.Peek()
}

func (s *MinStack2[T]) Min() (T, bool) {
	return s.min.Peek()
}

func (s *MinStack2[T]) IsEmpty() bool {
	return s.top.IsEmpty()
}
