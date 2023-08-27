package adjacencymatrix

import (
	"strconv"
)

const (
	MaxVertices = 50
)

type Graph struct {
	v   int      // number of vertices
	adj [][]bool // adjacency list
}

func CreateGraph(v int) *Graph {
	adj := make([][]bool, v)
	for i := 0; i < v; i++ {
		adj[i] = make([]bool, v)
	}
	g := &Graph{
		v:   v,
		adj: adj,
	}

	return g
}

func (g *Graph) AddEdge(v int, w int) {
	g.adj[v][w] = true
}

func (g *Graph) Bfs(s int) string {
	result := ""
	visited := make([]bool, MaxVertices)
	queue := make([]int, MaxVertices)
	front, rear := 0, 0

	visited[s] = true
	queue[rear] = s
	rear++

	for front != rear {
		s = queue[front]
		front++
		result += strconv.Itoa(s)
		for adjacent := 0; adjacent < g.v; adjacent++ {
			if g.adj[s][adjacent] && !visited[adjacent] {
				visited[adjacent] = true
				queue[rear] = adjacent
				rear++
			}
		}
		if front != rear {
			result += " "
		}
	}
	return result
}
