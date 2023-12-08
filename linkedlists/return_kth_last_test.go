package linkedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindKLastNode(t *testing.T) {
	list := createLetterList([]string{"A", "B", "C"})

	assert.Equal(t, "C", findKthLastNode(list, 0).data)
	assert.Equal(t, "B", findKthLastNode(list, 1).data)
	assert.Equal(t, "A", findKthLastNode(list, 2).data)
	assert.Nil(t, findKthLastNode(list, 3))
	assert.Nil(t, findKthLastNode(list, -1))
}

func TestFindKLastNodeRecursively(t *testing.T) {
	list := createLetterList([]string{"A", "B", "C"})

	assert.Equal(t, "C", findKthLastNodeRecursively(list, 0).data)
	assert.Equal(t, "B", findKthLastNodeRecursively(list, 1).data)
	assert.Equal(t, "A", findKthLastNodeRecursively(list, 2).data)
	assert.Nil(t, findKthLastNodeRecursively(list, 3))
	assert.Nil(t, findKthLastNodeRecursively(list, -1))
}
