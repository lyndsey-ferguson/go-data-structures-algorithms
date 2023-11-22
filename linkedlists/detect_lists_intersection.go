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
