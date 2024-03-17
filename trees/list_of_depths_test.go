package trees

import (
	"testing"

	"github.com/lyndsey-ferguson/go-data-structures-algorithms/linkedlists"
	"github.com/stretchr/testify/assert"
)

func TestListOfDepthsForUnbalancedTree(t *testing.T) {
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

	var list []*linkedlists.Node[int]
	list = ListOfDepths(tree, list, 0)
	assert.NotNil(t, list)
	assert.Equal(t, 4, len(list))
	assert.Equal(t, "13", list[0].ToString())
	assert.Equal(t, "6", list[1].ToString())
	assert.Equal(t, "2, 9", list[2].ToString())
	assert.Equal(t, "4, 10", list[3].ToString())
}

func TestListOfDepthsForBalancedTree(t *testing.T) {
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
	var list []*linkedlists.Node[int]
	list = ListOfDepths(tree, list, 0)
	assert.NotNil(t, list)
	assert.Equal(t, 4, len(list))
	assert.Equal(t, "13", list[0].ToString())
	assert.Equal(t, "6, 25", list[1].ToString())
	assert.Equal(t, "2, 9, 18, 26", list[2].ToString())
	assert.Equal(t, "4, 10, 19", list[3].ToString())
}
