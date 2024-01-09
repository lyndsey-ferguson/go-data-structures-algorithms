package linkedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStackSort(t *testing.T) {
	s := ComparableStack[int]{}
	s.Push(1)
	s.Push(8)
	s.Push(2)
	s.Push(5)
	s.Push(3)

	s.Sort()

	assert.Equal(t, "1, 2, 3, 5, 8", s.elements.ToString())

	u := ComparableStack[int]{}
	u.Push(3)
	u.Push(5)
	u.Push(2)
	u.Push(8)
	u.Push(1)

	u.Sort()
	assert.Equal(t, "1, 2, 3, 5, 8", u.elements.ToString())
}
