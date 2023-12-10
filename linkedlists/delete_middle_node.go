package linkedlists

/*
Implement an algorithm to delete a node in the middle
(i.e. any node but the first and last node, not necessarily the exact middle)
of a singly linked list, given only access to that node.

EXAMPLE:

Input the node c from the linked list a-> b -> c -> d -> e -> f
Result, nothing is returned, but the new list looks like a->b->d->e->f
*/
func deleteMiddleNode[T comparable](middle *Node[T]) {
	if middle == nil || middle.next == nil {
		return
	}
	middle.data = middle.next.data
	middle.next = middle.next.next
}
