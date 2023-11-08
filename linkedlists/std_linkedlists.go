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

/*
removeDuplicatesInline works, but let's review it for any bottlenecks, unnecessary
work, or duplicate work.
It is pretty straightforward, we're looping a smaller and smaller list for each
element in the array. The left hand side will be cleaned up, and the right hand
side will be the unknown --- at first.

However, we are duplicating work, we are reviewing each element in the list multiple
times. The worst case scenario is that there are no duplicates and each element
is reviewed. Not quite N^2 as we'll be visit N-1 nodes for the first node, N-2 nodes
for the second node, or basically the sum of N - 1 + N - 2 + N - 3 ... + 2 + 1,
which is n/2(n + 1), which is (N^2 + N)/2, which is basically O(N^2).
*/
func removeDuplicatesInline[T comparable](list *Node[T]) {
	if list == nil || list.next == nil {
		return
	}
	for cursor := list; cursor != nil; cursor = cursor.next {
		parent := cursor
		for runner := cursor.next; runner != nil; runner = runner.next {
			if runner.data == cursor.data {
				// remove the runner from the list
				// the parent's next no longer points to
				// the runner, so that, in go, effectively
				// will reove the runner Node
				parent.next = runner.next
			} else {
				// move the parent to point to the runner
				// as it is not a duplicate
				parent = runner
			}
		}
	}
}

/*
So, in the above algorithm, if we've visited a node once, we already "know" about
its value. Do we really need to look at it again?

No, we can track if we've already visited a node and, instead of multiple passes
of the same node, we can pass once. That would be a runtime of O(N)
*/

func removeDuplicatesWithHash[T comparable](list *Node[T]) {
	if list == nil || list.next == nil {
		// if there are fewer than 1 elements, return early as
		// there is no chance for duplicates
		return
	}
	end := list         // keep track of the end of the list
	cursor := list.next // point to the next element

	// Create a map of nodes that we've already found
	// if it is unique, append the found node and
	// track it.
	foundNodes := map[T]int{list.data: 1}

	for cursor != nil {
		_, ok := foundNodes[cursor.data]
		if !ok {
			foundNodes[cursor.data] = 1
			end.next = cursor
			end = cursor
		}
		cursor = cursor.next
		end.next = nil
	}
	end.next = nil
}

func findKthLastNode[T comparable](list *Node[T], k int) *Node[T] {
	if k < 0 {
		return nil
	}

	var kth *Node[T]
	for cursor := list; cursor != nil; cursor = cursor.next {
		if k == 0 {
			kth = list
		} else if k < 0 && kth.next != nil {
			kth = kth.next
		}
		k = k - 1
	}
	return kth
}

/*
	There is also a recursive version of 'findKthLastNode'
*/

func findKthLastNodeRecursively[T comparable](list *Node[T], k int) *Node[T] {
	localK := k
	return _findKthLastNodeRecursively[T](list, &localK)
}

func _findKthLastNodeRecursively[T comparable](list *Node[T], k *int) *Node[T] {
	if list == nil {
		return nil
	} else {
		kth := _findKthLastNodeRecursively[T](list.next, k)
		if kth != nil {
			return kth
		} else if *k == 0 {
			return list
		} else {
			*k = *k - 1
			return nil
		}
	}
}

func deleteMiddleNode[T comparable](middle *Node[T]) {
	if middle == nil || middle.next == nil {
		return
	}
	cursor := middle
	for scout := middle.next; scout != nil; scout = scout.next {
		cursor.data = scout.data
		if scout.next != nil {
			cursor = cursor.next
		}
	}
	cursor.next = nil
}
