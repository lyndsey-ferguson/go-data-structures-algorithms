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

type RemoveStrategy int8

const (
	RemoveFirstStrategy = iota
	RemoveAllStrategy
)

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

func RemoveNodesWithValueUsingStrategy(list *Node, data int, strategy RemoveStrategy) *Node {
	for list != nil && list.data == data {
		list = list.next
		if strategy == RemoveFirstStrategy {
			return list
		}
	}
	if list == nil {
		return list
	}

	for cursor, previous := list.next, list; cursor != nil; cursor, previous = cursor.next, cursor {
		if cursor.data == data {
			previous.next = cursor.next
			if strategy == RemoveFirstStrategy {
				break
			}
		}
	}
	return list
}

func RemoveFirstNodeWithValue(list *Node, data int) *Node {
	return RemoveNodesWithValueUsingStrategy(list, data, RemoveFirstStrategy)
}

func RemoveAllNodesWithValue(list *Node, data int) *Node {
	return RemoveNodesWithValueUsingStrategy(list, data, RemoveAllStrategy)
}
