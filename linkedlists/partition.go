package linkedlists

/*
Write code to partition a linked list around a value x,
such that all nodes less than x come before all nodes greater than or equal to x.
If x is contained within the list, the values of x only need to be after the elements
less than x (see below).
The partition element x can appear anywhere in the “right partition”: it does not
need to appear between the left and right partitions.

EXAMPLE:
Input 3 -> 5 -> 8 -> 5 -> 10 -> 2 -> 1 (partition = 5)
output: 3 -> 1 -> 2    ->    10 -> 5 -> 5 -> 8
*/

/*
The first intuition is to create two partitions, one a "low number" partition,
and a "equal/high number partition". Then iterate over the list, put the nodes
with a low number into the low-number-partition, and the other nodes into the
equal-high-number-partition.

My question is, is there an undue amount of emphasis placed on the unimportance
of the partitioned list not requiring any order, just be anywhere in the higher
number partition?

Maybe the idea is that I am to avoid making two partitions.

Ok, maybe I can iterate the list with two pointers, one "low" and one high"
Move the low pointer until it reaches a node with a value equal to or higher
and then swap the two values. Then, find the next node that is equal to or
higher, starting from the last found "low" pointer, and the
next high (or equal) value starting from the "high" pointer.

Swap swap swap

*/

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

/*
Just for fun, let's create the two lists, and then merge them
*/
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
