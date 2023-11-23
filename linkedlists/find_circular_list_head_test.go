package linkedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindCircularListHeadWithHash(t *testing.T) {
	list := createLetterList([]string{"A", "B", "C", "D", "E"})
	cNode := findNode[string](list, func(n *Node[string]) bool {
		return n.data == "C"
	})
	AppendNode(list, cNode)

	assert.Equal(t, cNode, findCircularListHeadWithHash(list))
}

func TestFindCircularListHeadWithHashReturnsNil(t *testing.T) {
	list := createLetterList([]string{"A", "B", "C", "D", "E"})

	assert.Nil(t, findCircularListHeadWithHash(list))
}
