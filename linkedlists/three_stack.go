package linkedlists

import (
	"errors"
	"fmt"
)

type ThreeStack[T comparable] struct {
	elements []T
	tops     []int
}

type IntThreeStack ThreeStack[uint32]

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

/*
I'm thinking that I can make a more compact three-stack, but it will be a little more costly
in terms of runtime.

If I can reserve 3 bits of an array element, I can store up to 29 bits for a positive number.
I could add more logic to support negative numbers, but for the time being, I will keep it
simple.

Now, when number is pushed, I look for the first available empty element and put the value there
with the stack number encoded in the top 3 bits. When a number is popped, I shift any stacks that
have their top on the right of that last number 1 element back and update the tops, and for the
stack that was popped, I update that top to the next element in that stack to the left.
*/

func makeIntThreeStack(capacity int) *IntThreeStack {
	ts := IntThreeStack{
		elements: make([]uint32, capacity),
		tops:     []int{-1, -1, -1},
	}
	return &ts
}

func _get28BitNumber(element uint32) uint32 {
	return (element & ^uint32(7<<29))
}

func (ts *IntThreeStack) Push(stack byte, value uint32) error {
	if value > uint32(1<<29) {
		return errors.New(fmt.Sprintf("Value %d too large for compact array", value))
	}
	top := ts.tops[stack-1]
	for i := top + 1; i < len(ts.elements); i++ {
		if ts.elements[i] == uint32(0) {
			ts.tops[stack-1] = i
			value |= uint32(1 << (28 + stack))
			ts.elements[i] = value
			return nil
		}
	}
	return errors.New("Stackoverflow")
}

func (ts *IntThreeStack) IsEmpty(stack byte) bool {
	return ts.tops[stack-1] == -1
}

func (ts *IntThreeStack) Peek(stack byte) (uint32, error) {
	if ts.IsEmpty(stack) {
		return 0, errors.New("Stack Underflow")
	}
	top := ts.tops[stack-1]
	return _get28BitNumber(ts.elements[top]), nil
}

func (ts *IntThreeStack) _shiftOtherStacksLeft(notStack byte, top int) {
	// for each other stack, if their top is greater
	// than the index of the "popped" element, reduce
	// that value by 1
	for s := 0; s < 3; s++ {
		if s != int(notStack-1) {
			t := ts.tops[s]
			if t > top {
				ts.tops[s] -= 1
			}
		}
	}
	// shift all elements after the "popped" element
	// to the left
	e := top
	for ; e+1 < len(ts.elements); e++ {
		if ts.elements[e+1] == 0 {
			ts.elements[e] = 0
		} else {
			ts.elements[e] = ts.elements[e+1]
		}
	}
	// find the previous element that is associated with this stack
	ts.elements[e] = 0
}

func (ts *IntThreeStack) _findPreviousTopForStack(stack byte, top int) {
	b := uint32(1 << (28 + stack))
	for top = top - 1; top >= 0; top-- {
		value := ts.elements[top]
		if value&b > 0 {
			break
		}
	}
	ts.tops[stack-1] = top
}

func (ts *IntThreeStack) Pop(stack byte) (uint32, error) {
	if ts.IsEmpty(stack) {
		return 0, errors.New("Stack Underflow")
	}

	top := ts.tops[stack-1]
	result := ts.elements[top]
	ts._shiftOtherStacksLeft(stack, top)
	ts._findPreviousTopForStack(stack, top)

	return _get28BitNumber(result), nil
}
