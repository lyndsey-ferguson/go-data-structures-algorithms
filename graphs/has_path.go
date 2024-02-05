package graphs

import "github.com/lyndsey-ferguson/go-data-structures-algorithms/linkedlists"

func (g *ListGraph[T]) HasPath(a T, b T) bool {
	if len(g.nodes) < 2 {
		return false
	}
	toVisitQueue := linkedlists.Queue[*Node[T]]{}
	visited := make(map[T]bool)

	node := g.FindNode(a)
	toVisitQueue.Add(node)
	for r, ok := toVisitQueue.Remove(); ok; r, ok = toVisitQueue.Remove() {
		visited[r.value] = true
		if visited[a] && visited[b] {
			return true
		}
		for _, n := range r.children {
			if !visited[n.value] {
				toVisitQueue.Add(n)
			}
		}
	}
	return false
}
