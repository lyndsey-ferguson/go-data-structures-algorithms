package trees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNextInOrderTraversal(t *testing.T) {
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
