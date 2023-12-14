package linkedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeThreeStackHappyPath(t *testing.T) {
	ts := makeThreeStack[int](5)
	assert.NotNil(t, ts)

	assert.Equal(t, 15, len(ts.elements))
}

func TestIsEmptyThreeStackAfterCreation(t *testing.T) {
	ts := makeThreeStack[int](5)

	assert.True(t, ts.IsEmpty(1))
	assert.True(t, ts.IsEmpty(2))
	assert.True(t, ts.IsEmpty(3))
}

func TestPushThreeStackNotEmpty(t *testing.T) {
	ts := makeThreeStack[int](5)

	err := ts.Push(1, 1)
	assert.Nil(t, err)
	assert.False(t, ts.IsEmpty(1))
	assert.True(t, ts.IsEmpty(2))
	assert.True(t, ts.IsEmpty(3))

	err = ts.Push(2, 5)
	assert.Nil(t, err)
	assert.False(t, ts.IsEmpty(1))
	assert.False(t, ts.IsEmpty(2))
	assert.True(t, ts.IsEmpty(3))

	err = ts.Push(3, 10)
	assert.Nil(t, err)
	assert.False(t, ts.IsEmpty(1))
	assert.False(t, ts.IsEmpty(2))
	assert.False(t, ts.IsEmpty(3))
}

func TestPushThreeStackOverflows(t *testing.T) {
	ts := makeThreeStack[int](5)

	for i := 0; i < 4; i++ {
		err := ts.Push(1, i)
		assert.Nil(t, err)
		err = ts.Push(2, i+5)
		assert.Nil(t, err)
		err = ts.Push(3, i+10)
		assert.Nil(t, err)
	}
	err := ts.Push(1, 99)
	assert.Nil(t, err)
	err = ts.Push(2, 199)
	assert.Nil(t, err)
	err = ts.Push(3, 299)
	assert.Nil(t, err)
}

func TestPeekThreeStack(t *testing.T) {
	ts := makeThreeStack[int](5)

	ts.Push(1, 1)
	v, err := ts.Peek(1)
	assert.Nil(t, err)
	assert.Equal(t, 1, v)

	ts.Push(2, 5)
	v, err = ts.Peek(1)
	assert.Nil(t, err)
	assert.Equal(t, 1, v)
	v, err = ts.Peek(2)
	assert.Nil(t, err)
	assert.Equal(t, 5, v)

	ts.Push(3, 10)
	v, err = ts.Peek(1)
	assert.Nil(t, err)
	assert.Equal(t, 1, v)
	v, err = ts.Peek(2)
	assert.Nil(t, err)
	assert.Equal(t, 5, v)
	v, err = ts.Peek(3)
	assert.Nil(t, err)
	assert.Equal(t, 10, v)
}

func TestPeekEmptyThreeStack(t *testing.T) {
	ts := makeThreeStack[int](5)

	_, err := ts.Peek(1)
	assert.NotNil(t, err)
	_, err = ts.Peek(2)
	assert.NotNil(t, err)
	_, err = ts.Peek(3)
	assert.NotNil(t, err)
}

func TestPopThreeStack(t *testing.T) {
	ts := makeThreeStack[int](5)

	ts.Push(1, 1)
	ts.Push(2, 5)
	ts.Push(3, 10)

	v, err := ts.Pop(1)
	assert.Nil(t, err)
	assert.Equal(t, 1, v)
	_, err = ts.Pop(1)
	assert.NotNil(t, err)

	v, err = ts.Pop(3)
	assert.Nil(t, err)
	assert.Equal(t, 10, v)
	_, err = ts.Pop(3)
	assert.NotNil(t, err)

	v, err = ts.Pop(2)
	assert.Nil(t, err)
	assert.Equal(t, 5, v)
	_, err = ts.Pop(2)
	assert.NotNil(t, err)
}

func TestPopEmptyThreeStack(t *testing.T) {
	ts := makeThreeStack[int](5)

	_, err := ts.Pop(3)
	assert.NotNil(t, err)

	_, err = ts.Pop(2)
	assert.NotNil(t, err)

	_, err = ts.Pop(1)
	assert.NotNil(t, err)
}

func TestIsEmptyEmptyThreeStack(t *testing.T) {
	ts := makeThreeStack[int](5)

	assert.True(t, ts.IsEmpty(1))
	assert.True(t, ts.IsEmpty(2))
	assert.True(t, ts.IsEmpty(3))
}
