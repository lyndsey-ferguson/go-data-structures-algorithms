package linkedlists

import (
	"fmt"
)

type Node[T comparable] struct {
	data T
	next *Node[T]
}

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

func (s *Stack[T]) Print() {
	fmt.Printf("%s\n", s.elements.ToString())
}

func (list *Node[T]) ToString() string {
	var result string
	for cursor := list; cursor != nil; {
		result = result + fmt.Sprintf("%v", cursor.data)
		cursor = cursor.next
		if cursor != nil {
			result = result + ", "
		}
	}
	return result
}

func CreateNode[T comparable](data T) *Node[T] {
	node := new(Node[T])
	node.data = data
	return node
}

func AppendNode[T comparable](list *Node[T], node *Node[T]) *Node[T] {
	if list == nil {
		return node
	}
	cursor := list
	for cursor.next != nil {
		cursor = cursor.next
	}
	cursor.next = node
	return list
}

func InsertAfter[T comparable](nodeBefore *Node[T], nodeAfter *Node[T]) {
	if nodeBefore != nil {
		nodeAfter.next = nodeBefore.next
		nodeBefore.next = nodeAfter
	}
}

func RemoveAfter[T comparable](node *Node[T]) {
	if node != nil && node.next != nil {
		node.next = node.next.next
	}
}

func SearchNodeIteratively[T comparable](list *Node[T], data T) *Node[T] {
	for cursor := list; cursor != nil; {
		if cursor.data == data {
			return cursor
		}
		cursor = cursor.next
	}
	return nil
}

func SearchNodeRecursively[T comparable](list *Node[T], data T) *Node[T] {
	if list == nil {
		return nil
	}
	if list.data == data {
		return list
	}
	return SearchNodeRecursively(list.next, data)
}

type RemoveStrategy int8

const (
	RemoveFirstStrategy = iota
	RemoveAllStrategy
)

func ReverseList[T comparable](list *Node[T]) *Node[T] {
	var reversedList *Node[T]
	for cursor := list; cursor != nil; {
		tmpNext := cursor.next
		cursor.next = reversedList
		reversedList = cursor
		cursor = tmpNext
	}
	return reversedList
}

func RemoveNodesWithValueUsingStrategy[T comparable](list *Node[T], data T, strategy RemoveStrategy) *Node[T] {
	for list != nil && list.data == data {
		list = list.next
		if strategy == RemoveFirstStrategy {
			return list
		}
	}
	if list == nil {
		return list
	}

	for cursor, previous := list.next, list; cursor != nil; cursor, previous = cursor.next, cursor {
		if cursor.data == data {
			previous.next = cursor.next
			if strategy == RemoveFirstStrategy {
				break
			}
		}
	}
	return list
}

func RemoveFirstNodeWithValue[T comparable](list *Node[T], data T) *Node[T] {
	return RemoveNodesWithValueUsingStrategy(list, data, RemoveFirstStrategy)
}

func RemoveAllNodesWithValue[T comparable](list *Node[T], data T) *Node[T] {
	return RemoveNodesWithValueUsingStrategy(list, data, RemoveAllStrategy)
}
