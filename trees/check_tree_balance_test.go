package trees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBalancedTreeHeightPasses(t *testing.T) {
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
	height, balanced := tree.GetBalancedTreeHeight()
	assert.True(t, balanced)
	assert.Equal(t, 4, height)
}

func TestGetBalancedTreeHeightFails(t *testing.T) {
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
	)
	height, balanced := tree.GetBalancedTreeHeight()
	assert.False(t, balanced)
	assert.Equal(t, 4, height)
}
