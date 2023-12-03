package linkedlists

import (
	"fmt"
)

type Stack[T comparable] struct {
	elements *Node[T]
}

func (s *Stack[T]) Push(i T) {
	n := CreateNode(i)

	if s.elements == nil {
		s.elements = n
	} else {
		n.next = s.elements
		s.elements = n
	}
}

func (s *Stack[T]) IsEmpty() bool {
	return s.elements == nil
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.elements == nil {
		var nonExistant T
		return nonExistant, false
	}
	i := s.elements.data
	s.elements = s.elements.next

	return i, true
}

func (s *Stack[T]) Peek() (T, bool) {
	if s.elements == nil {
		var nonExistant T
		return nonExistant, false
	}
	return s.elements.data, true
}

func (s *Stack[T]) Print() {
	fmt.Printf("%s\n", s.elements.ToString())
}
