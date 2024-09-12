package linkedlists

func SwapNodesInPairs[T comparable](head *Node[T]) (result *Node[T]) {
	// Create a temporary node to serve as a previous node
	// in front of head
	if head == nil || head.next == nil {
		return head
	}
	var noValue T
	dummy := CreateNode(noValue)
	dummy.next = head

	previous := dummy

	first := head       // first = -> 1 -> ...
	second := head.next // second = -> 2 -> ...

	// dummy = -> <0> -> 1 -> 2 -> 3 -> 4 -> 5
	for second != nil && first != second {
		// previous = -> 1 -> 3 -> 4 -> 5 -> NULL
		previous.next = second
		// previous = -> 1 -> 5 -> NULL
		first.next = second.next
		// first = -> 3 -> 5 -> NULL
		second.next = first
		// second = -> 4 -> 3 -> 5 -> NULL

		previous = first   // previous = -> 3 -> 5 -> NULL
		first = first.next // first = -> 5 -> NULL
		if first != nil {
			second = first.next // second = NULL
		} else {
			second = nil
		}
	}
	return dummy.next
}
