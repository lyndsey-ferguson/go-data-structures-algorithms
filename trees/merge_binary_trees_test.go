package trees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeBinaryTrees(t *testing.T) {
	t1 := CreateBinaryTree([]int{1, 2, 3})
	t2 := CreateBinaryTree([]int{1, 2, 3})

	t3 := CreateBinaryTree([]int{2, 4, 6})

	assert.True(t, true, t3.Equal(mergeTrees(t1, t2)))

	t1 = CreateBinaryTree([]int{2, 1})
	t2 = CreateBinaryTree([]int{2, 3})

	t3 = CreateBinaryTree([]int{4, 4})
	assert.True(t, true, t3.Equal(mergeTrees(t1, t2)))

	t1 = CreateBinaryTree([]int{3, 4, 5})
	t2 = CreateBinaryTree([]int{1})

	t3 = CreateBinaryTree([]int{4, 4, 5})
	assert.True(t, true, t3.Equal(mergeTrees(t1, t2)))
}
