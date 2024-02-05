package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBfsStartingAtVertix2(t *testing.T) {
	g := CreateMatrixGraph(4)

	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	g.AddEdge(2, 0)
	g.AddEdge(2, 3)
	g.AddEdge(3, 3)

	actualVisitedStr := g.Bfs(2)
	assert.Equal(t, "2 0 3 1", actualVisitedStr)
}

func TestDfsStartingAtVertix0(t *testing.T) {
	g := CreateMatrixGraph(5)

	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(0, 3)
	g.AddEdge(1, 0)
	g.AddEdge(2, 0)
	g.AddEdge(2, 3)
	g.AddEdge(2, 4)
	g.AddEdge(3, 0)
	g.AddEdge(3, 2)
	g.AddEdge(4, 2)

	actualVisitedStr := g.Dfs(0)
	assert.Equal(t, "0 3 2 4 1", actualVisitedStr)
}

func TestDfsStartingAtVertix2(t *testing.T) {
	g := CreateMatrixGraph(4)

	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	g.AddEdge(2, 0)
	g.AddEdge(2, 1)
	g.AddEdge(2, 3)
	g.AddEdge(3, 3)

	actualVisitedStr := g.Dfs(2)
	assert.Equal(t, "2 3 1 0", actualVisitedStr)
}
