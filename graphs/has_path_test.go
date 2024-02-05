package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasPath(t *testing.T) {
	g := ListGraph[string]{}
	g.AddEdge("A", "B")
	g.AddEdge("B", "A")
	g.AddEdge("A", "C")
	g.AddEdge("C", "D")
	g.AddEdge("D", "E")
	g.AddEdge("D", "F")
	assert.True(t, g.HasPath("A", "E"))
	assert.False(t, g.HasPath("E", "F"))
}
