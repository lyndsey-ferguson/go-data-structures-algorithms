package linkedlists

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func convertLinkedListToString(list *Node) string {
	var result string
	for cursor := list; cursor != nil; {
		result = result + strconv.Itoa(cursor.data)
		cursor = cursor.next
		if cursor != nil {
			result = result + ", "
		}
	}
	return result
}

func TestBasicNodeCreation(t *testing.T) {
	node := CreateNode(1)
	assert.Equal(t, 1, node.data)
	assert.Equal(t, "1", convertLinkedListToString(node))
}

func TestAppendNodeToList(t *testing.T) {
	list := CreateNode(1)
	node := CreateNode(2)
	list = AppendNode(list, node)
	assert.Equal(t, node, list.next)
	assert.Equal(t, "1, 2", convertLinkedListToString(list))
}

func TestAppendNodeToEmptyList(t *testing.T) {
	node := CreateNode(2)
	list := AppendNode(nil, node)
	assert.Equal(t, node, list)
	assert.Equal(t, "2", convertLinkedListToString(list))
}

func TestAppendToTwoElementList(t *testing.T) {
	list := CreateNode(1)
	node2 := CreateNode(2)
	node3 := CreateNode(3)
	list = AppendNode(list, node2)
	list = AppendNode(list, node3)

	assert.Equal(t, node2, list.next)
	assert.Equal(t, node3, list.next.next)
	assert.Equal(t, "1, 2, 3", convertLinkedListToString(list))
}

func TestInsertAfterNode(t *testing.T) {
	list := CreateNode(1)
	node2 := CreateNode(2)
	node3 := CreateNode(3)
	list = AppendNode(list, node2)
	list = AppendNode(list, node3)
	node4 := CreateNode(4)

	InsertAfter(node2, node4)

	assert.Equal(t, "1, 2, 4, 3", convertLinkedListToString(list))
}

func TestRemoveAfter(t *testing.T) {
	list := CreateNode(1)
	node2 := CreateNode(2)
	node3 := CreateNode(3)
	list = AppendNode(list, node2)
	list = AppendNode(list, node3)

	RemoveAfter(list)
	assert.Equal(t, "1, 3", convertLinkedListToString(list))
}

// taken from
func TestSearchNodeIteratively(t *testing.T) {
	list := CreateNode(1)
	node2 := CreateNode(2)
	node3 := CreateNode(3)
	list = AppendNode(list, node2)
	list = AppendNode(list, node3)

	foundNode := SearchNodeIteratively(list, 2)
	assert.Equal(t, node2, foundNode)
	assert.Equal(t, 2, foundNode.data)

	foundNode = SearchNodeIteratively(list, 12)
	assert.Nil(t, foundNode)
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

func TestReverseList(t *testing.T) {
	list := CreateNode(1)
	node2 := CreateNode(2)
	node3 := CreateNode(3)
	list = AppendNode(list, node2)
	list = AppendNode(list, node3)

	list = ReverseList(list)
	assert.Equal(t, "3, 2, 1", convertLinkedListToString(list))
}
