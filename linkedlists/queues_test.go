package linkedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	q := Queue[int]{}

	q.Add(1)
}

func TestQueueRemoveNothing(t *testing.T) {
	q := Queue[int]{}

	v, ok := q.Remove()
	assert.Zero(t, v)
	assert.False(t, ok)
}

func TestQueueRemoveLastValue(t *testing.T) {
	q := Queue[int]{}

	q.Add(9)

	v, ok := q.Remove()
	assert.Equal(t, v, 9)
	assert.True(t, ok)
}

func TestQueueRemoveIsNothingAfterLastValue(t *testing.T) {
	q := Queue[int]{}

	q.Add(9)

	q.Remove()
	v, ok := q.Remove()
	assert.Zero(t, v)
	assert.False(t, ok)
}

func TestQueueRemoveValuesInCorrectOrder(t *testing.T) {
	q := Queue[int]{}

	q.Add(9)
	q.Add(101)

	v, _ := q.Remove()
	assert.Equal(t, v, 9)
	v, _ = q.Remove()
	assert.Equal(t, v, 101)
}

func TestQueueEmptyInitially(t *testing.T) {
	q := Queue[int]{}

	assert.True(t, q.IsEmpty())
}

func TestQueueEmptyAfterRemovingLastItem(t *testing.T) {
	q := Queue[int]{}
	q.Add(9)
	q.Remove()
	assert.True(t, q.IsEmpty())
}

func TestQueueNotEmpty(t *testing.T) {
	q := Queue[int]{}
	q.Add(8)
	assert.False(t, q.IsEmpty())
}

func TestQueuePeekReturnsNothing(t *testing.T) {
	q := Queue[int]{}
	v, ok := q.Peek()
	assert.False(t, ok)
	assert.Zero(t, v)
}

func TestQueuePeekReturnsFirstQueuedElement(t *testing.T) {
	q := Queue[int]{}
	q.Add(9)
	q.Add(101)

	v, ok := q.Peek()
	assert.Equal(t, v, 9)
	assert.True(t, ok)
}

func TestQueuePeekDoesNotRemove(t *testing.T) {
	q := Queue[int]{}
	q.Add(101)

	q.Peek()
	assert.False(t, q.IsEmpty())
}
