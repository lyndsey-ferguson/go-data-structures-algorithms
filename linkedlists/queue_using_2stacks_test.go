package linkedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyQueueEnqueue(t *testing.T) {
	q := MyQueue[int]{}

	q.Enqueue(7)
}

func TestMyQueueDequeue(t *testing.T) {
	q := MyQueue[int]{}

	v, ok := q.Dequeue()
	assert.False(t, ok)

	q.Enqueue(7)
	q.Enqueue(9)

	v, ok = q.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, 7, v)

	v, ok = q.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, 9, v)

	v, ok = q.Dequeue()
	assert.False(t, ok)
}

func TestMyQueuePeek(t *testing.T) {
	q := MyQueue[int]{}

	v, ok := q.Peek()
	assert.False(t, ok)

	q.Enqueue(7)
	q.Enqueue(9)

	v, ok = q.Peek()
	assert.True(t, ok)
	assert.Equal(t, 7, v)

	v, ok = q.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, 7, v)

	v, ok = q.Peek()
	assert.True(t, ok)
	assert.Equal(t, 9, v)

	q.Dequeue()
	v, ok = q.Peek()
	assert.False(t, ok)
}
