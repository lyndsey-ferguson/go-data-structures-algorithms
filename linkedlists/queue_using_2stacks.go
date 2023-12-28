package linkedlists

/*
Implement q MyQueue class which implements a queue using two stacks

A queue is a first-in, first-out, structure. This challenge is asking
us to make a queue using 2 stacks, which are first-in, last-out data
structures.

We need the following methods:
1. enqueue, adds an element to the end of the queue
2. dequeue, removes an element from the front of the queue
3. peek, returns the value of the element at the front of the queue

Maybe pushes go to one stack, and a pop pops all items from the first
stack to the second stack, and the last item on the second stack.

It feels wasteful to keep moving elements back and forth. Well, let's
write it up for now and then see if something better pops out at me.
*/

type MyQueue[T comparable] struct {
	pushes      Stack[T]
	pops        Stack[T]
	peekedValue T
}

func (q *MyQueue[T]) Enqueue(v T) {
	for popped, ok := q.pops.Pop(); ok; popped, ok = q.pops.Pop() {
		q.pushes.Push(popped)
	}
	if q.pushes.IsEmpty() {
		q.peekedValue = v
	}
	q.pushes.Push(v)
}

func (q *MyQueue[T]) Dequeue() (T, bool) {
	for pushed, ok := q.pushes.Pop(); ok; pushed, ok = q.pushes.Pop() {
		q.pops.Push(pushed)
	}
	v, ok := q.pops.Pop()
	if !q.pops.IsEmpty() {
		q.peekedValue, _ = q.pops.Peek()
	}
	return v, ok
}

func (q *MyQueue[T]) Peek() (T, bool) {
	if !(q.pushes.IsEmpty() && q.pops.IsEmpty()) {
		return q.peekedValue, true
	}
	return q.peekedValue, false
}
