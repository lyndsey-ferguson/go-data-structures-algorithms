package linkedlists

func partitionIntList[T int32](list *Node[T], partition T) {
	if list == nil {
		// there are no elements
		return
	}

	findLowlyPlacedHighNodeComparator := func(node *Node[T]) bool {
		return (node != nil && node.data >= partition)
	}
	findHighlyPlacedLowNodeComparator := func(node *Node[T]) bool {
		return (node != nil && node.data < partition)
	}
	// Find the first value that is greater than, or equal
	// to the partition value
	low := findNode(list, findLowlyPlacedHighNodeComparator)
	if low == nil {
		return
	}
	// Find the next value that is less than the partition value
	// So we can swap
	high := findNode(low.next, findHighlyPlacedLowNodeComparator)
	for low != nil && high != nil {
		low.data, high.data = high.data, low.data
		low = findNode(low.next, findLowlyPlacedHighNodeComparator)
		high = findNode(high.next, findHighlyPlacedLowNodeComparator)
	}
}

func partitionIntListMerging[T int32](list **Node[T], parition T) {
	if list == nil || *list == nil {
		return
	}
	var smallerList *Node[T]
	var equalOrGreaterList *Node[T]

	for cursor := *list; cursor != nil; {
		if cursor.data < parition {
			if smallerList == nil {
				smallerList = cursor
			} else {
				AppendNode(smallerList, cursor) // We could make this faster by tracking the tail of the smaller list
			}
		} else {
			if equalOrGreaterList == nil {
				equalOrGreaterList = cursor
			} else {
				AppendNode(equalOrGreaterList, cursor) // We could make this faster by tracking the tail of the greater list
			}
		}
		tail := cursor
		cursor = cursor.next
		tail.next = nil
	}
	if smallerList != nil {
		AppendNode(smallerList, equalOrGreaterList) // We could make this faster by tracking the tail of the smaller list
	} else {
		smallerList = equalOrGreaterList
	}
	*list = smallerList
}
