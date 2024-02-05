package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func AssertChildrenFound[T comparable](t *testing.T, expectedChildren []T, actualChildren []*Node[T]) {
	foundChild := false
	for _, expectedChild := range expectedChildren {
		for _, childNode := range actualChildren {
			if childNode.value == expectedChild {
				foundChild = true
				break
			}
		}
		assert.True(t, foundChild, "Expected child node '%s' not found", expectedChild)
	}
}

func TestAddEdge(t *testing.T) {
	g := ListGraph[string]{}
	g.AddEdge("A", "B")
	g.AddEdge("B", "A")
	g.AddEdge("A", "C")
	g.AddEdge("C", "D")
	g.AddEdge("D", "E")
	g.AddEdge("D", "F")

	assert.Equal(t, 6, len(g.nodes))
	expectedNodeValues := []string{"A", "B", "C", "D", "E", "F"}
	foundNodes := make(map[string]*Node[string])
	for _, node := range g.nodes {
		foundNodes[node.value] = node
	}
	for _, expectedNodeValue := range expectedNodeValues {
		_, foundNode := foundNodes[expectedNodeValue]
		assert.True(t, foundNode, "Expected node '%s' not found", expectedNodeValue)
	}
	n := foundNodes["A"]
	AssertChildrenFound(t, []string{"B", "C"}, n.children)
	n = foundNodes["B"]
	AssertChildrenFound(t, []string{"A"}, n.children)
	n = foundNodes["C"]
	AssertChildrenFound(t, []string{"D"}, n.children)
	n = foundNodes["D"]
	AssertChildrenFound(t, []string{"E"}, n.children)
	n = foundNodes["E"]
	assert.Zero(t, len(n.children))
	n = foundNodes["F"]
	assert.Zero(t, len(n.children))
}
