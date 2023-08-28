package linkedlists

import (
	"fmt"
	"strconv"
)

type Node struct {
	data int
	next *Node
}

type Stack struct {
	elements *Node
}

func (s *Stack) Push(i int) {
	n := CreateNode(i)

	if s.elements == nil {
		s.elements = n
	} else {
		n.next = s.elements
		s.elements = n
	}
}

func (s *Stack) IsEmpty() bool {
	return s.elements == nil
}

func (s *Stack) Pop() (int, bool) {
	if s.elements == nil {
		return 0, false
	}
	i := s.elements.data
	s.elements = s.elements.next

	return i, true
}

func (s *Stack) Print() {
	fmt.Printf("%s\n", s.elements.ToString())
}

func (list *Node) ToString() string {
	var result string
	for cursor := list; cursor != nil; {
		result = result + strconv.Itoa(cursor.data)
		cursor = cursor.next
		if cursor != nil {
			result = result + ", "
		}
	}
	return result
}

func CreateNode(data int) *Node {
	node := new(Node)
	node.data = data
	return node
}

func AppendNode(list *Node, node *Node) *Node {
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

func InsertAfter(nodeBefore *Node, nodeAfter *Node) {
	if nodeBefore != nil {
		nodeAfter.next = nodeBefore.next
		nodeBefore.next = nodeAfter
	}
}

func RemoveAfter(node *Node) {
	if node != nil && node.next != nil {
		node.next = node.next.next
	}
}

func SearchNodeIteratively(list *Node, data int) *Node {
	for cursor := list; cursor != nil; {
		if cursor.data == data {
			return cursor
		}
		cursor = cursor.next
	}
	return nil
}

func SearchNodeRecursively(list *Node, data int) *Node {
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

func ReverseList(list *Node) *Node {
	var reversedList *Node
	for cursor := list; cursor != nil; {
		tmpNext := cursor.next
		cursor.next = reversedList
		reversedList = cursor
		cursor = tmpNext
	}
	return reversedList
}

func RemoveNodesWithValueUsingStrategy(list *Node, data int, strategy RemoveStrategy) *Node {
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

func RemoveFirstNodeWithValue(list *Node, data int) *Node {
	return RemoveNodesWithValueUsingStrategy(list, data, RemoveFirstStrategy)
}

func RemoveAllNodesWithValue(list *Node, data int) *Node {
	return RemoveNodesWithValueUsingStrategy(list, data, RemoveAllStrategy)
}
