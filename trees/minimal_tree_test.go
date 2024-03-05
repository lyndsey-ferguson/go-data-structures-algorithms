package trees

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTreeifySortedArray(t *testing.T) {
	a := []int{2, 4, 6, 9, 10, 13, 18, 19, 25, 26}
	tree := TreeifySortedArray(a, 0, len(a)-1)

	var b bytes.Buffer
	tree.Print(&b)

	expectedTree := CreateNode(13,
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
	treesEqual := tree.Equal(expectedTree)
	assert.True(t, treesEqual, "Generated tree does not match expectated tree")
	if !treesEqual {
		fmt.Println("Generated Tree:")
		tree.Print(os.Stdout)
		fmt.Println("\nExpected Tree:")
		expectedTree.Print(os.Stdout)
	}
}
