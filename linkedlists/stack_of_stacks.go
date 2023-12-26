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
