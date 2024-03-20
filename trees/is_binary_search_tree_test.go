package trees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsBinarySearchTreeForNonBst(t *testing.T) {
	tree := CreateNode(13,
		CreateNode(6,
			CreateNode(12,
				nil,
				CreateNode(41),
			),
			CreateNode(9,
				nil,
				CreateNode(10),
			),
		),
		CreateNode(25,
			CreateNode(18,
				nil,
				CreateNode(19),
			),
			CreateNode(26),
		),
	)
	assert.False(t, IsBinarySearchTree(tree))
}

func TestIsBinarySearchTreeForBst(t *testing.T) {
	tree := CreateNode(13,
		CreateNode(6,
			CreateNode(2,
				nil,
				CreateNode(4),
			),
			CreateNode(9,
				nil,
				CreateNode(10),
			),
		),
		CreateNode(25,
			CreateNode(18,
				nil,
				CreateNode(19),
			),
			CreateNode(26),
		),
	)
	assert.True(t, IsBinarySearchTree(tree))
}

func TestIsBinarySearchTreeForFakeBst(t *testing.T) {
	tree := CreateNode(20,
		CreateNode(10,
			nil,
			CreateNode(25),
		),
		CreateNode(30),
	)
	assert.False(t, IsBinarySearchTree(tree))
}
