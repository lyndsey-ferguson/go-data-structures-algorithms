package linkedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSwapNodesInPairsHappyPath(t *testing.T) {
	twoPairsList := createList([]int{1, 2, 3, 4})
	expectedResult := createList([]int{2, 1, 4, 3})

	twoPairsList = SwapNodesInPairs(twoPairsList)

	assert.Equal(t, expectedResult, twoPairsList)
	twoPairsAndOneList := createList([]int{1, 2, 3, 4, 5})
	expectedResult = createList([]int{2, 1, 4, 3, 5})

	twoPairsAndOneList = SwapNodesInPairs(twoPairsAndOneList)
	assert.Equal(t, expectedResult, twoPairsAndOneList)
}

func TestSwapNodesInPairsEdgeCases(t *testing.T) {
	oneNodeList := CreateNode(1)
	assert.Equal(t, oneNodeList, SwapNodesInPairs(oneNodeList))

	var nilList *Node[int]
	assert.Nil(t, SwapNodesInPairs(nilList))

	oneNodeList.next = oneNodeList
	assert.Equal(t, oneNodeList, SwapNodesInPairs(oneNodeList))
}
