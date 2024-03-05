package trees

import (
	"fmt"
	"io"
	"os"
)

type Node[T comparable] struct {
	data  T
	left  *Node[T]
	right *Node[T]
}

func CreateNode[T comparable](data T, childNodes ...*Node[T]) *Node[T] {
	n := Node[T]{
		data: data,
	}
	if childrenCount := len(childNodes); childrenCount > 0 {
		n.left = childNodes[0]
		if childrenCount > 1 {
			n.right = childNodes[1]
		}
	}
	return &n
}

func (tree *Node[T]) Equal(otherTree *Node[T]) bool {
	if tree == nil && otherTree == nil {
		return true
	} else if tree == nil || otherTree == nil {
		return false
	}
	if tree.data == otherTree.data {
		return tree.left.Equal(otherTree.left) && tree.right.Equal(otherTree.right)
	}
	return false
}

func (tree *Node[T]) PrintHelper(space int, out io.Writer) {
	// Base case
	if tree == nil {
		return
	}

	// Increase distance between levels
	space += 10

	// Process right child first
	tree.right.PrintHelper(space, out)

	// Print current node after space count
	fmt.Fprintln(out)
	for i := 10; i < space; i++ {
		fmt.Fprint(out, " ")
	}
	fmt.Fprintf(out, "%v\n", tree.data)

	// Process left child
	tree.left.PrintHelper(space, out)
}

func (tree *Node[T]) Print(out io.Writer) {
	if out == nil {
		out = os.Stdout
	}
	tree.PrintHelper(0, out)
}
