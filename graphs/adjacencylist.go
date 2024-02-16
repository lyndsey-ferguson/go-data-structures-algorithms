package graphs

type Node[T comparable] struct {
	value    T
	children []*Node[T]
}

type ListGraph[T comparable] struct {
	nodes []*Node[T]
}

/* Add edge to graph */
/* Adding an edge adds the vertex too */
func (g *ListGraph[T]) AddEdge(u T, v T) {
	var nonExistentValue T
	uNode := &Node[T]{}
	vNode := &Node[T]{}

	foundNodes := 0
	for _, node := range g.nodes {
		if node.value == u {
			uNode = node
			foundNodes += 1
		} else if node.value == v {
			vNode = node
			foundNodes += 1
		}
		if foundNodes > 1 {
			break
		}
	}
	if uNode.value == nonExistentValue {
		uNode = &Node[T]{
			value: u,
		}
		g.nodes = append(g.nodes, uNode)
	}
	if vNode.value == nonExistentValue {
		vNode = &Node[T]{
			value: v,
		}
		g.nodes = append(g.nodes, vNode)
	}
	uNode.children = append(uNode.children, vNode)
}

func (g *ListGraph[T]) FindNode(u T) *Node[T] {
	for _, node := range g.nodes {
		if node.value == u {
			return node
		}
	}
	return nil
}
