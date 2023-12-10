package linkedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteMiddleNode(t *testing.T) {
	list := createLetterList([]string{"A", "B", "C", "D", "E", "F"})
	expectedList := createLetterList([]string{"A", "B", "D", "E", "F"})
	middle := list.next.next

	assert.Equal(t, "C", middle.data)
	deleteMiddleNode(middle)
	assert.Equal(t, expectedList.ToString(), list.ToString())

	expectedList2 := createLetterList([]string{"A", "D", "E", "F"})
	deleteMiddleNode(list.next)
	assert.Equal(t, expectedList2.ToString(), list.ToString())
}
