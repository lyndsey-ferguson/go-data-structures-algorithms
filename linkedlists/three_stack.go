package linkedlists

import "errors"

type ThreeStack[T comparable] struct {
	elements []T
	tops     []int
}

func makeThreeStack[T comparable](capacity int) *ThreeStack[T] {
	ts := ThreeStack[T]{
		elements: make([]T, capacity*3),
		tops:     []int{-1, capacity - 1, (capacity * 2) - 1},
	}
	return &ts
}

func (ts *ThreeStack[T]) _getStackIndices(stack int) (int, int) {
	stackSize := len(ts.elements) / 3
	startIndex := ((stack - 1) * stackSize)
	endIndex := (stack * stackSize) - 1

	return startIndex, endIndex
}

func (ts *ThreeStack[T]) Push(stack int, value T) error {
	_, endIndex := ts._getStackIndices(stack)
	top := ts.tops[stack-1]
	if top < endIndex {
		top += 1
		ts.tops[stack-1] += 1
	} else {
		return errors.New("Stack Overflow!")
	}
	ts.elements[top] = value
	return nil
}

func (ts *ThreeStack[T]) Peek(stack int) (T, error) {
	startIndex, endIndex := ts._getStackIndices(stack)
	top := ts.tops[stack-1]
	if top >= startIndex && top <= endIndex {
		return ts.elements[top], nil
	}
	var none T
	return none, errors.New("Stack Underflow!")
}

func (ts *ThreeStack[T]) Pop(stack int) (T, error) {
	startIndex, endIndex := ts._getStackIndices(stack)
	top := ts.tops[stack-1]
	if top >= startIndex && top <= endIndex {
		result := ts.elements[top]
		ts.tops[stack-1] -= 1
		return result, nil
	}
	var none T
	return none, errors.New("Stack Underflow!")
}

func (ts *ThreeStack[T]) IsEmpty(stack int) bool {
	startIndex, _ := ts._getStackIndices(stack)
	top := ts.tops[stack-1]
	hasElements := (top >= startIndex)
	return !hasElements
}
