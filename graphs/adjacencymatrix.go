package graphs

import (
	"strconv"

	"github.com/lyndsey-ferguson/go-data-structures-algorithms/linkedlists"
)

const (
	MaxVertices = 50
)

type MatrixGraph struct {
	v   int      // number of vertices
	adj [][]bool // adjacency list
}

func CreateMatrixGraph(v int) *MatrixGraph {
	adj := make([][]bool, v)
	for i := 0; i < v; i++ {
		adj[i] = make([]bool, v)
	}
	g := &MatrixGraph{
		v:   v,
		adj: adj,
	}

	return g
}

func (g *MatrixGraph) AddEdge(v int, w int) {
	g.adj[v][w] = true
}

func (g *MatrixGraph) Bfs(s int) string {
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

func (g *MatrixGraph) dfs(s int, visited []bool) string {
	var stack linkedlists.Stack[int]
	var result string

	stack.Push(s)
	for !stack.IsEmpty() {
		top, _ := stack.Pop()
		if visited[top] {
			continue
		} else {
			visited[top] = true
			if len(result) > 0 {
				result += " "
			}
			result += strconv.Itoa(top)
		}
		for adjacent := 0; adjacent < g.v; adjacent++ {
			if g.adj[top][adjacent] {
				if !visited[adjacent] {
					stack.Push(adjacent)
				}
			}
		}
	}
	return result
}
func (g *MatrixGraph) Dfs(s int) string {
	visited := make([]bool, MaxVertices)
	return g.dfs(s, visited)
}
