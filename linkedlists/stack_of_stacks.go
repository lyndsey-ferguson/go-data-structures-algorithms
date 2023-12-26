package linkedlists

type StackOfStacks[T comparable] struct {
	capacity int
	consumed int
	stacks   *Node[Stack[T]]
}

func (ss *StackOfStacks[T]) Push(v T) {
	if ss.consumed >= ss.capacity || ss.stacks == nil {
		s := Stack[T]{}
		sn := CreateNode[Stack[T]](s)
		sn.next = ss.stacks
		ss.stacks = sn
		ss.consumed = 0
	}
	ss.stacks.data.Push(v)
	ss.consumed++
}

func (ss *StackOfStacks[T]) Pop() (T, bool) {
	if ss.stacks == nil {
		var nonExistant T
		return nonExistant, false
	}
	v, ok := ss.stacks.data.Pop()
	if ok && ss.stacks.data.IsEmpty() {
		ss.stacks = ss.stacks.next
	}
	return v, ok
}

func (ss *StackOfStacks[T]) Peek() (T, bool) {
	if ss.stacks == nil {
		var nonExistant T
		return nonExistant, false
	}
	return ss.stacks.data.Peek()
}

func (ss *StackOfStacks[T]) IsEmpty() bool {
	return ss.stacks == nil
}

func (ss *StackOfStacks[T]) PopAt(stack int) (T, bool) {
	// for the sake of understanding, let's get the count of stacks
	stackCount := 0
	for cursor := ss.stacks; cursor != nil; cursor, stackCount = cursor.next, stackCount+1 {
	}

	// if stackCount == stack, then we can just pop from the first stack
	if stack == stackCount-1 {
		v, ok := ss.stacks.data.Pop()
		if ok && ss.stacks.data.IsEmpty() {
			ss.stacks = ss.stacks.next
		}
		return v, ok
	}
	// otherwise, let's iterate until we find the stackIndex that matches the given stack
	// 0 is for the last stack, 1 is for the 2nd last stack, 2, is for the 3rd last stack
	// and N-1 is for the second stack

	parent := ss.stacks
	cursor := ss.stacks.next
	index := stackCount - 2
	for ; index != stack && cursor != nil; index, cursor = index-1, cursor.next {
		parent = parent.next
	}
	if index == stack && cursor != nil {
		v, ok := cursor.data.Pop()
		if ok && cursor.data.IsEmpty() {
			parent.next = cursor.next
		}
		return v, ok
	}
	var nonExistant T
	return nonExistant, false
}
