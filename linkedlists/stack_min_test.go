package linkedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinStackPush(t *testing.T) {
	s := MinStack[int]{}

	s.Push(1)
}

func TestMinStackPopNothing(t *testing.T) {
	s := MinStack[int]{}

	v, ok := s.Pop()
	assert.Zero(t, v)
	assert.False(t, ok)
}

func TestMinStackPopLastValue(t *testing.T) {
	s := MinStack[int]{}

	s.Push(9)

	v, ok := s.Pop()
	assert.Equal(t, v, 9)
	assert.True(t, ok)
}

func TestMinStackPopIsNothingAfterLastValue(t *testing.T) {
	s := MinStack[int]{}

	s.Push(9)

	s.Pop()
	v, ok := s.Pop()
	assert.Zero(t, v)
	assert.False(t, ok)
}

func TestMinStackPopsValuesInCorrectOrder(t *testing.T) {
	s := MinStack[int]{}

	s.Push(9)
	s.Push(101)

	v, _ := s.Pop()
	assert.Equal(t, v, 101)
	v, _ = s.Pop()
	assert.Equal(t, v, 9)
}

func TestMinStackEmptyInitially(t *testing.T) {
	s := MinStack[int]{}

	assert.True(t, s.IsEmpty())
}

func TestMinStackEmptyAfterPoppingLastItem(t *testing.T) {
	s := MinStack[int]{}
	s.Push(9)
	s.Pop()
	assert.True(t, s.IsEmpty())
}

func TestMinStackNotEmpty(t *testing.T) {
	s := MinStack[int]{}
	s.Push(8)
	assert.False(t, s.IsEmpty())
}

func TestMinStackPeekReturnsNothing(t *testing.T) {
	s := MinStack[int]{}
	v, ok := s.Peek()
	assert.False(t, ok)
	assert.Zero(t, v)
}

func TestMinStackPeekReturnsTopElement(t *testing.T) {
	s := MinStack[int]{}
	s.Push(9)
	s.Push(101)

	v, ok := s.Peek()
	assert.Equal(t, v, 101)
	assert.True(t, ok)
}

func TestMinStackPeekDoesNotPop(t *testing.T) {
	s := MinStack[int]{}
	s.Push(101)

	s.Peek()
	assert.False(t, s.IsEmpty())
}

func TestMinStackMinWorksWithEmpty(t *testing.T) {
	s := MinStack[int]{}

	_, ok := s.Min()
	assert.False(t, ok)
}

func TestMinStackMinWithOneElement(t *testing.T) {
	s := MinStack[int]{}

	s.Push(1)

	v, ok := s.Min()
	assert.True(t, ok)
	assert.Equal(t, 1, v)
}

func TestMinStackMinWithManyElements(t *testing.T) {
	s := MinStack[int]{}

	s.Push(10)
	s.Push(11)
	s.Push(9)
	s.Push(7)
	s.Push(12)
	s.Push(8)
	s.Push(13)
	s.Push(91)

	v, ok := s.Min()
	assert.True(t, ok)
	assert.Equal(t, 7, v)
}

func TestMinStackMinWithManyElementsAfterPops(t *testing.T) {
	s := MinStack[int]{}

	s.Push(10)
	s.Push(11)
	s.Push(9)
	s.Push(7)
	s.Push(12)
	s.Push(8)
	s.Push(13)
	s.Push(91)

	for i := 0; i < 5; i++ {
		s.Pop()
	}
	v, ok := s.Min()
	assert.True(t, ok)
	assert.Equal(t, 9, v)
}
