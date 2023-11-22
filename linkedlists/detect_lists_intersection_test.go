package linkedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntersectionNodeForListsUsingMap(t *testing.T) {
	list1 := createLetterList([]string{"3", "1", "5", "9", "7", "2", "1"})
	list2 := createLetterList([]string{"4", "6"})

	expectedNode := SearchNodeIteratively(list1, "7")
	sixNode := SearchNodeIteratively(list2, "6")
	sixNode.next = expectedNode

	assert.Equal(t, expectedNode, intersectionNodeForListsUsingMap(list1, list2))
}
