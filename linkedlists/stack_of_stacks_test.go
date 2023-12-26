package linkedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStackOfStacksPush(t *testing.T) {
	ss := StackOfStacks[int]{
		capacity: 3,
	}

	ss.Push(1)
	ss.Push(2)
	ss.Push(3)
	ss.Push(4)
	ss.Push(5)
}

func TestStackOfStacksPop(t *testing.T) {
	ss := StackOfStacks[int]{
		capacity: 3,
	}

	_, ok := ss.Pop()
	assert.False(t, ok)
	ss.Push(1)
	ss.Push(2)
	ss.Push(3)
	ss.Push(4)
	ss.Push(5)

	for i := 5; i >= 1; i-- {
		v, ok := ss.Pop()
		assert.True(t, ok)
		assert.Equal(t, i, v)
	}
}

func TestStackOfStacksPeek(t *testing.T) {
	ss := StackOfStacks[int]{
		capacity: 3,
	}
	_, ok := ss.Peek()
	assert.False(t, ok)
	ss.Push(1)
	ss.Push(2)
	ss.Push(3)
	ss.Push(4)
	ss.Push(5)
	v, ok := ss.Peek()
	assert.True(t, ok)
	assert.Equal(t, 5, v)

	for i := 5; i >= 1; i-- {
		ss.Pop()
	}
	_, ok = ss.Peek()
	assert.False(t, ok)
}

func TestStackOfStacksIsEmpty(t *testing.T) {
	ss := StackOfStacks[int]{
		capacity: 3,
	}
	assert.True(t, ss.IsEmpty())
	ss.Push(1)
	ss.Push(2)
	ss.Push(3)
	ss.Push(4)
	ss.Push(5)
	assert.False(t, ss.IsEmpty())

	for i := 5; i >= 1; i-- {
		v, ok := ss.Pop()
		assert.True(t, ok)
		assert.Equal(t, i, v)
	}
	assert.True(t, ss.IsEmpty())
}
