package general

import (
	"container/list"
)

func findMaxElementInKayWindows(arr []int, k int) []int {
	if len(arr) < 1 {
		return []int{}
	}
	if k > len(arr) {
		k = len(arr)
	}

	result := []int{}

	d := list.New()

	// iterate i over k elements in the list
	//   if the value of the ith element is larger than
	//   the element at the rear of the deque, then remove
	//.  the rear element
	// when the k-window is processed, increment the window
	// also pop any elements from the deque if their indice
	// is before the window

	for i := 0; i < len(arr)-k+1; i = i + 1 {
		for e := d.Front(); e != nil; e = d.Front() {
			if frontIndex, ok := e.Value.(int); ok {
				if frontIndex < i {
					d.Remove(e)
				} else {
					break
				}
			}
		}
		for w := i; w < i+k; w = w + 1 {
			for e := d.Back(); e != nil; e = d.Back() {
				if backIndex, ok := e.Value.(int); ok && arr[backIndex] < arr[w] {
					d.Remove(e)
				} else {
					break
				}
			}
			d.PushBack(w)
		}
		e := d.Front()
		if e != nil {
			if frontIndex, ok := e.Value.(int); ok {
				result = append(result, arr[frontIndex])
			}
		}
	}
	return result
}
