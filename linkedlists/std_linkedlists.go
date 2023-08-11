package linkedlists

type Node struct {
	data int
	next *Node
}

func CreateNode(data int) *Node {
	node := new(Node)
	node.data = data
	return node
}

func AppendNode(list *Node, node *Node) *Node {
	if list == nil {
		return node
	}
	cursor := list
	for cursor.next != nil {
		cursor = cursor.next
	}
	cursor.next = node
	return list
}

func InsertAfter(nodeBefore *Node, nodeAfter *Node) {
	if nodeBefore != nil {
		nodeAfter.next = nodeBefore.next
		nodeBefore.next = nodeAfter
	}
}

func RemoveAfter(node *Node) {
	if node != nil && node.next != nil {
		node.next = node.next.next
	}
}

func SearchNodeIteratively(list *Node, data int) *Node {
	for cursor := list; cursor != nil; {
		if cursor.data == data {
			return cursor
		}
		cursor = cursor.next
	}
	return nil
}

func SearchNodeRecursively(list *Node, data int) *Node {
	if list == nil {
		return nil
	}
	if list.data == data {
		return list
	}
	return SearchNodeRecursively(list.next, data)
}

func ReverseList(list *Node) *Node {
	var reversedList *Node
	for cursor := list; cursor != nil; {
		tmpNext := cursor.next
		cursor.next = reversedList
		reversedList = cursor
		cursor = tmpNext
	}
	return reversedList
}
