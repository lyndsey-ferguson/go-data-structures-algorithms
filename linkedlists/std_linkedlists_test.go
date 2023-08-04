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
}

func TestAppendNodeToEmptyList(t *testing.T) {
	node := CreateNode(2)
	list := AppendNode(nil, node)
	if list != node {
		t.Errorf("list = %x; want %x", &list.next, &node)
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
}
