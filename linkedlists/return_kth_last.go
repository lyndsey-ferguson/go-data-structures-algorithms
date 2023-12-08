package linkedlists

/*
Implement an algorithm to find the kth to last element on a singly linked list
*/

/*
Let's walk through some examples

Let's say we have this list:

list -> A -> B -> C

with k = 0, return c
with k = 1, return b,
with k = 2, return a,
with k = 3, return nil

For this algorithm, I'd say, let's start iterating on the list. As we go down the list, whenever k = 0
we can assign k to the list

# For example

list -> A -> B -> C
k = 0

when we start looping with 'cursor', and k = 0, we point kth to list

cursor -> A
kth -> A

# As cursor moves forward, as long as the next node is not nil, kth moves forward

cursor -> B
kth -> B

cursor -> C
kth -> C

cursor -> nil
exit with kth -> C

--

If k = 1, then we want to wait until k = 0 to assign kth to the list

cursor -> A
k = 1
kth -> nil

cursor -> B
k = 0
kth -> A

cursor -> C
k = -1
kth -> B

cursor -> nil
exit with kth -> B
*/
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
	There is also a recursive version of 'findKthLastNode'. We can think of this as a 'stack' which
	is what recursion uses. For every element, we put it on the stack, when we reach the end, we
	pop off that number of elements and that is our answer.
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
