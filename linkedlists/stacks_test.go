package linkedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStackPush(t *testing.T) {
	s := Stack[int]{}

	s.Push(1)
}

func TestStackPopNothing(t *testing.T) {
	s := Stack[int]{}

	v, ok := s.Pop()
	assert.Zero(t, v)
	assert.False(t, ok)
}

func TestStackPopLastValue(t *testing.T) {
	s := Stack[int]{}

	s.Push(9)

	v, ok := s.Pop()
	assert.Equal(t, v, 9)
	assert.True(t, ok)
}

func TestStackPopIsNothingAfterLastValue(t *testing.T) {
	s := Stack[int]{}

	s.Push(9)

	s.Pop()
	v, ok := s.Pop()
	assert.Zero(t, v)
	assert.False(t, ok)
}

func TestStackPopsValuesInCorrectOrder(t *testing.T) {
	s := Stack[int]{}

	s.Push(9)
	s.Push(101)

	v, _ := s.Pop()
	assert.Equal(t, v, 101)
	v, _ = s.Pop()
	assert.Equal(t, v, 9)
}

func TestStackEmptyInitially(t *testing.T) {
	s := Stack[int]{}

	assert.True(t, s.IsEmpty())
}

func TestStackEmptyAfterPoppingLastItem(t *testing.T) {
	s := Stack[int]{}
	s.Push(9)
	s.Pop()
	assert.True(t, s.IsEmpty())
}

func TestStackNotEmpty(t *testing.T) {
	s := Stack[int]{}
	s.Push(8)
	assert.False(t, s.IsEmpty())
}

func TestStackPeekReturnsNothing(t *testing.T) {
	s := Stack[int]{}
	v, ok := s.Peek()
	assert.False(t, ok)
	assert.Zero(t, v)
}

func TestStackPeekReturnsTopElement(t *testing.T) {
	s := Stack[int]{}
	s.Push(9)
	s.Push(101)

	v, ok := s.Peek()
	assert.Equal(t, v, 101)
	assert.True(t, ok)
}

func TestStackPeekDoesNotPop(t *testing.T) {
	s := Stack[int]{}
	s.Push(101)

	s.Peek()
	assert.False(t, s.IsEmpty())
}
