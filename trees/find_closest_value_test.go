package trees

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindClosestValue(t *testing.T) {
	a := []int{4, 2, 6, 1, 5, 7}
	slices.Sort(a)
	root := CreateBinaryTree(a)

	assert.Equal(t, 4, findClosestValue(root, 3.5))

	a = []int{8, 3, 10, 1, 6, 14}
	slices.Sort(a)

	root = CreateBinaryTree(a)
	assert.Equal(t, 3, findClosestValue(root, 4.4))

	a = []int{5, 3, 8, 1, 4, 6, 9}
	root = CreateBinaryTree(a)
	assert.Equal(t, 8, findClosestValue(root, 7.4))
}
