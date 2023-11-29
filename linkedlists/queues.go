package linkedlists

import (
	"fmt"
)

type Queue[T comparable] struct {
	first *Node[T]
	last  *Node[T]
}

func (q *Queue[T]) Add(i T) {
	n := CreateNode(i)

	if q.last != nil {
		q.last.next = n
	}
	q.last = n
	if q.first == nil {
		q.first = q.last
	}
}

func (q *Queue[T]) IsEmpty() bool {
	return q.first == nil
}

func (q *Queue[T]) Remove() (T, bool) {
	if q.first == nil {
		var nonExistant T
		return nonExistant, false
	}
	data := q.first.data
	q.first = q.first.next
	if q.first == nil {
		q.last = nil
	}
	return data, true
}

func (q *Queue[T]) Peek() (T, bool) {
	if q.first == nil {
		var nonExistant T
		return nonExistant, false
	}
	return q.first.data, true
}

func (q *Queue[T]) Print() {
	fmt.Printf("%s\n", q.first.ToString())
}
