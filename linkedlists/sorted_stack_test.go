package linkedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortedStackPushAndPop(t *testing.T) {
	s := SortedStack[int]{}

	s.Push(3)
	s.Push(6)
	s.Push(4)
	s.Push(5)
	s.Push(1)
	s.Push(0)
	s.Push(7)
	s.Push(2)

	v, ok := s.Pop()
	for i := 0; i < 8 && ok; i = i + 1 {
		assert.True(t, ok)
		assert.Equal(t, i, v)
		v, ok = s.Pop()
	}
}
