package linkedlists

/*
Given two (singly) linked lists, determine if the two lists intersect.
Return the intersection node.
Note that the intersection is defined based on reference, not value.
That is if the kth element of the first linked list is the exact same node
(by references) as the jth node of the second linked list, then they are intersecting.


For example:

b -> c -> d -> e
     ^
	 |
	 o
	 ^
	 |
	 d

The intersection is 'c'

I could use a hashmap of pointers to find the node that appears twice.
*/

func intersectionNodeForListsUsingMap[T comparable](list1 *Node[T], list2 *Node[T]) *Node[T] {
	foundNodes := make(map[*Node[T]]bool)

	for cursor := list1; cursor != nil; cursor = cursor.next {
		foundNodes[cursor] = true
	}
	for cursor := list2; cursor != nil; cursor = cursor.next {
		_, ok := foundNodes[cursor]
		if ok {
			return cursor
		} else {
			foundNodes[cursor] = true
		}
	}

	return nil
}

/*
However, there is something unique about this case: both lists
have the same final, tail, node. If they don't have the same
tail, then we know these lists do not intersect.

Could I get the length of both lists, and then, starting at
the same number of nodes away from the end of the longest
list, and then step forward until we find the intersection
node.
*/
func getIthNode[T comparable](list *Node[T], i int) *Node[T] {
	start := list
	for count := 0; count < i && start != nil; count += 1 {
		start = start.next
	}
	return start
}

func intersectionNodeForListsUsingLength[T comparable](list1 *Node[T], list2 *Node[T]) *Node[T] {
	var tail *Node[T]

	list1Length := 0
	list2Length := 0

	for cursor := list1; cursor != nil; cursor = cursor.next {
		list1Length += 1
		if cursor.next == nil {
			tail = cursor
		}
	}
	for cursor := list2; cursor != nil; cursor = cursor.next {
		list2Length += 1
		if cursor.next == nil && cursor != tail {
			// We've reached the tail of list2, but it is NOT the tail of list1, these lists
			// do not intersect
			return nil
		}
	}
	start1 := getIthNode(list1, list1Length-list2Length)
	start2 := getIthNode(list2, list2Length-list1Length)
	for ; start1 != nil && start2 != nil; start1, start2 = start1.next, start2.next {
		if start1 == start2 {
			return start1
		}
	}
	return nil
}
