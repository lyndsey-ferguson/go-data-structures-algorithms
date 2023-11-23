package linkedlists

func findCircularListHeadWithHash[T comparable](list *Node[T]) *Node[T] {
	visitedNodes := make(map[*Node[T]]bool)

	for cursor := list; cursor != nil; cursor = cursor.next {
		_, ok := visitedNodes[cursor]
		if ok {
			return cursor
		}
		visitedNodes[cursor] = true
	}
	return nil
}
