package trees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateBinaryTree(t *testing.T) {
	a := []int{4, 2, 6, 1, 5, 7}

	root := CreateBinaryTree(a)
	assert.True(t, IsBinarySearchTree(root))

	a = []int{8, 3, 10, 1, 6, 14}
	root = CreateBinaryTree(a)
	assert.True(t, IsBinarySearchTree(root))
}
