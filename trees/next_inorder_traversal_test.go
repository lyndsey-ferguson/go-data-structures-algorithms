package trees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNextInOrderTraversalSimpleCase(t *testing.T) {
	root := CreateParentedNode(1)
	left := CreateParentedNode(0)
	left.parent = root
	root.left = left
	right := CreateParentedNode(2)
	right.parent = root
	root.right = right

	assert.Equal(t, right, NextInOrderTraversal(root))
	assert.Equal(t, root, NextInOrderTraversal(left))

	var unparentedNode *ParentedNode[int]
	assert.Equal(t, unparentedNode, NextInOrderTraversal(right))
}

func TestNextInOrderTraversalBiggerTree(t *testing.T) {
	gNode := CreateParentedNode('g')
	cNode := CreateParentedNode('C')
	tree := CreateParentedNode('D',
		CreateParentedNode('B',
			CreateParentedNode('A'),
			cNode,
		),
		CreateParentedNode('F',
			CreateParentedNode('E'),
			gNode,
		),
	)
	var unparentedNode *ParentedNode[rune]
	assert.Equal(t, unparentedNode, NextInOrderTraversal(gNode))

	assert.Equal(t, tree, NextInOrderTraversal(cNode))
}
