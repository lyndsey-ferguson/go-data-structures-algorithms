package adjacencymatrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBfsStartingAtVertix2(t *testing.T) {
	g := CreateGraph(4)

	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	g.AddEdge(2, 0)
	g.AddEdge(2, 3)
	g.AddEdge(3, 3)

	actualVisitedStr := g.Bfs(2)
	assert.Equal(t, "2 0 3 1", actualVisitedStr)
}
