package linkedlists

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
