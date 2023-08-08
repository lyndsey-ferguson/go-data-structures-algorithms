package main

import (
	"testing"
)

func TestBasicNodeCreation(t *testing.T) {
	node := CreateNode(1)
	if node.data != 1 {
		t.Errorf("node.data = %d; want 1", node.data)
	}
}

func TestAppendNodeToList(t *testing.T) {
	list := CreateNode(1)
	node := CreateNode(2)
	list = AppendNode(list, node)
	if list.next != node {
		t.Errorf("list.next = %x; want %x", &list.next, &node)
	}
	if list.data != 1 {
		t.Errorf("list.data = %x; want %x", list.data, 1)
	}
	if list.next.data != 2 {
		t.Errorf("list.next.data = %x; want %x", list.next.data, 2)
	}
}

func TestAppendNodeToEmptyList(t *testing.T) {
	node := CreateNode(2)
	list := AppendNode(nil, node)
	if list != node {
		t.Errorf("list = %x; want %x", &list.next, &node)
	}
	if list.data != 2 {
		t.Errorf("list.data = %x; want %x", list.data, 2)
	}
}

func TestAppendToTwoElementList(t *testing.T) {
	list := CreateNode(1)
	node2 := CreateNode(2)
	node3 := CreateNode(3)
	list = AppendNode(list, node2)
	list = AppendNode(list, node3)

	if list.next != node2 {
		t.Errorf("list.next = %x; want %x", &list.next, &node2)
	}
	if list.next.next != node3 {
		t.Errorf("list.next.next = %x; want %x", &list.next, &node3)
	}
	if list.data != 1 {
		t.Errorf("list.data = %x; want %x", list.data, 1)
	}
	if list.next.data != 2 {
		t.Errorf("list.data = %x; want %x", list.data, 2)
	}
	if list.next.next.data != 3 {
		t.Errorf("list.data = %x; want %x", list.data, 3)
	}
}

func TestInsertAfterNode(t *testing.T) {
	list := CreateNode(1)
	node2 := CreateNode(2)
	node3 := CreateNode(3)
	list = AppendNode(list, node2)
	list = AppendNode(list, node3)
	node4 := CreateNode(4)

	InsertAfter(node2, node4)
	if list.data != 1 {
		t.Errorf("list.data = %x; want %x", list.data, 1)
	}
	if list.next.data != 2 {
		t.Errorf("list.next.data = %x; want %x", list.data, 2)
	}
	if list.next.next.data != 4 {
		t.Errorf("list.next.next.data = %x; want %x", list.data, 4)
	}
	if list.next.next.next.data != 3 {
		t.Errorf("list.next.next.next.data = %x; want %x", list.data, 3)
	}
}

func TestRemoveAfter(t *testing.T) {
	list := CreateNode(1)
	node2 := CreateNode(2)
	node3 := CreateNode(3)
	list = AppendNode(list, node2)
	list = AppendNode(list, node3)

	RemoveAfter(list)
	if list.data != 1 {
		t.Errorf("list.data = %d; want %d", list.data, 1)
	}
	if list.next.data != 3 {
		t.Errorf("list.next.data = %d; want %d", list.next.data, 3)
	}
}

// taken from
func TestSearchNodeIteratively(t *testing.T) {
	list := CreateNode(1)
	node2 := CreateNode(2)
	node3 := CreateNode(3)
	list = AppendNode(list, node2)
	list = AppendNode(list, node3)

	foundNode := SearchNodeIteratively(list, 2)
	if foundNode.data != 2 {
		t.Errorf("foundNode.data = %d; want %d", foundNode.data, 2)
	}
	foundNode = SearchNodeIteratively(list, 12)
	if foundNode != nil {
		t.Errorf("foundNode is %x; want nil", &foundNode)
	}
}

func TestSearchNodeRecursively(t *testing.T) {
	list := CreateNode(1)
	node2 := CreateNode(2)
	node3 := CreateNode(3)
	list = AppendNode(list, node2)
	list = AppendNode(list, node3)

	foundNode := SearchNodeRecursively(list, 2)
	if foundNode.data != 2 {
		t.Errorf("foundNode.data = %d; want %d", foundNode.data, 2)
	}
	foundNode = SearchNodeRecursively(list, 12)
	if foundNode != nil {
		t.Errorf("foundNode is %x; want nil", &foundNode)
	}
}
