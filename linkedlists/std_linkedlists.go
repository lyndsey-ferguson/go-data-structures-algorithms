package main

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
