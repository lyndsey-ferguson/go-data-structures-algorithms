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

	list = createLetterList([]string{"A", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L"})
	dNode := findNode[string](list, func(n *Node[string]) bool {
		return n.data == "D"
	})
	AppendNode(list, dNode)
	assert.Equal(t, dNode, findCircularListHeadWithHash(list))
}

func TestFindCircularListHeadWithHashReturnsNil(t *testing.T) {
	list := createLetterList([]string{"A", "B", "C", "D", "E"})

	assert.Nil(t, findCircularListHeadWithHash(list))
}

func TestFindCircularListHeadWithRunners(t *testing.T) {
	list := createLetterList([]string{"A", "B", "C", "D", "E"})
	cNode := findNode[string](list, func(n *Node[string]) bool {
		return n.data == "C"
	})
	AppendNode(list, cNode)

	assert.Equal(t, cNode, findCircularListHeadWithRunners(list))

	list = createLetterList([]string{"A", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L"})
	dNode := findNode[string](list, func(n *Node[string]) bool {
		return n.data == "D"
	})
	AppendNode(list, dNode)
	assert.Equal(t, dNode, findCircularListHeadWithRunners(list))
}

func TestFindCircularListHeadWithRunnersReturnsNil(t *testing.T) {
	list := createLetterList([]string{"A", "B", "C", "D", "E"})

	assert.Nil(t, findCircularListHeadWithRunners(list))
}
