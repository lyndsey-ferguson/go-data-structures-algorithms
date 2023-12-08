package linkedlists

/*
Remove duplicates from an unsorted linked list
*/

/*
A duplicate is a copy of another node's value.
The 1st node that is found is the 'original' and can stay.

"Unsorted" seems important, why mention that? Maybe indicating that we cannot just
get a node's value and then look for that in the next node(s) to remove dups.

So, the rough way to do this may bbe to have two pointers?

For each value, run ahead and look for it's duplicate. When the dup is found, remove it.
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
				// will remove the runner Node
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
