package linkedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicNodeCreation(t *testing.T) {
	node := CreateNode(1)
	assert.Equal(t, 1, node.data)
	assert.Equal(t, "1", node.ToString())
}

func TestAppendNodeToList(t *testing.T) {
	list := CreateNode(1)
	node := CreateNode(2)
	list = AppendNode(list, node)
	assert.Equal(t, node, list.next)
	assert.Equal(t, "1, 2", list.ToString())
}

func TestAppendNodeToEmptyList(t *testing.T) {
	node := CreateNode(2)
	list := AppendNode(nil, node)
	assert.Equal(t, node, list)
	assert.Equal(t, "2", list.ToString())
}

func TestAppendToTwoElementList(t *testing.T) {
	list := CreateNode(1)
	node2 := CreateNode(2)
	node3 := CreateNode(3)
	list = AppendNode(list, node2)
	list = AppendNode(list, node3)

	assert.Equal(t, node2, list.next)
	assert.Equal(t, node3, list.next.next)
	assert.Equal(t, "1, 2, 3", list.ToString())
}

func TestInsertAfterNode(t *testing.T) {
	list := CreateNode(1)
	node2 := CreateNode(2)
	node3 := CreateNode(3)
	list = AppendNode(list, node2)
	list = AppendNode(list, node3)
	node4 := CreateNode(4)

	InsertAfter(node2, node4)

	assert.Equal(t, "1, 2, 4, 3", list.ToString())
}

func TestRemoveAfter(t *testing.T) {
	list := CreateNode(1)
	node2 := CreateNode(2)
	node3 := CreateNode(3)
	list = AppendNode(list, node2)
	list = AppendNode(list, node3)

	RemoveAfter(list)
	assert.Equal(t, "1, 3", list.ToString())
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
	assert.Equal(t, "3, 2, 1", list.ToString())
}

func TestRemoveFirstNodeWithValue(t *testing.T) {
	list := CreateNode(1)
	node2 := CreateNode(2)
	node3 := CreateNode(3)
	list = AppendNode(list, node2)
	list = AppendNode(list, node3)

	updatedList := RemoveFirstNodeWithValue(list, 2)
	assert.Equal(t, "1, 3", updatedList.ToString())

	list2 := CreateNode(0)
	node4 := CreateNode(4)
	node5 := CreateNode(5)
	list2 = AppendNode(list2, node4)
	list2 = AppendNode(list2, node5)
	updatedList = RemoveFirstNodeWithValue(list2, 0)
	assert.Equal(t, "4, 5", updatedList.ToString())

	updatedList = RemoveFirstNodeWithValue(list2, 9)
	assert.Equal(t, "0, 4, 5", updatedList.ToString())

	updatedList = RemoveFirstNodeWithValue(nil, 2)
	assert.Equal(t, "", updatedList.ToString())
}

func TestRemoveAllNodesWithValue(t *testing.T) {
	list := CreateNode(2)
	node2 := CreateNode(2)
	node3 := CreateNode(2)
	list = AppendNode(list, node2)
	list = AppendNode(list, node3)

	updatedList := RemoveAllNodesWithValue(list, 2)
	assert.Equal(t, "", updatedList.ToString())

	list2 := CreateNode(0)
	node4 := CreateNode(4)
	node5 := CreateNode(0)
	list2 = AppendNode(list2, node4)
	list2 = AppendNode(list2, node5)
	updatedList = RemoveAllNodesWithValue(list2, 0)
	assert.Equal(t, "4", updatedList.ToString())

	updatedList = RemoveAllNodesWithValue(updatedList, 9)
	assert.Equal(t, "4", updatedList.ToString())

	updatedList = RemoveAllNodesWithValue(nil, 2)
	assert.Equal(t, "", updatedList.ToString())
}

func createLetterList(letters []string) *Node[string] {
	if len(letters) < 1 {
		return nil
	}
	list := CreateNode(letters[0])

	for i, end := 1, list; i < len(letters); i, end = i+1, end.next {
		node := CreateNode(letters[i])
		AppendNode(end, node)
	}
	return list
}

func TestRemoveDuplicatesInline(t *testing.T) {
	list := createLetterList([]string{"x", "w", "y", "x", "h", "y", "a", "x"})
	expectedList := createLetterList([]string{"x", "w", "y", "h", "a"})

	removeDuplicatesInline(list)
	assert.Equal(t, expectedList.ToString(), list.ToString())
}

func TestRemoveDuplicatesWithHash(t *testing.T) {
	list := createLetterList([]string{"x", "w", "y", "x", "h", "y", "a", "x"})
	expectedList := createLetterList([]string{"x", "w", "y", "h", "a"})

	removeDuplicatesWithHash(list)
	assert.Equal(t, expectedList.ToString(), list.ToString())
}

func TestFindKLastNode(t *testing.T) {
	list := createLetterList([]string{"A", "B", "C"})

	assert.Equal(t, "C", findKthLastNode(list, 0).data)
	assert.Equal(t, "B", findKthLastNode(list, 1).data)
	assert.Equal(t, "A", findKthLastNode(list, 2).data)
	assert.Nil(t, findKthLastNode(list, 3))
	assert.Nil(t, findKthLastNode(list, -1))
}
