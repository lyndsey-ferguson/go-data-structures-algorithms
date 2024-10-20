package heaps

import (
	"container/heap"
)

// An StringHeap is a min-heap of Strings.
type StringHeap []string

func (h StringHeap) Len() int           { return len(h) }
func (h StringHeap) Less(i, j int) bool { return len(h[i]) > len(h[j]) }
func (h StringHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *StringHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(string))
}

func (h *StringHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *StringHeap) Stringify() string {
	result := ""
	for h.Len() > 0 {
		result += heap.Pop(h).(string)
	}
	return result
}

func frequencySort(str string) string {
	// ToDo: Write Your Code Here.
	// I am considering two items

	// iterate through string to find character count of
	// each character
	// hash = [
	//      s : 4,
	//      e : 3,
	//      r : 2
	//      t : 1
	// ]
	//
	// once we get the hash, we can already see the pattern
	// if we could quickly sort this hash, that would be
	// best; however, what if we could create a hash of
	// these entities, then we could just recreate the
	// string ... or, we could create substrings and
	// add them. So that would be a string hash

	hash := make(map[rune]int)
	for _, r := range str {
		hash[r]++
	}

	h := &StringHeap{}
	heap.Init(h)

	for c, count := range hash {
		substr := make([]rune, count)
		for i := 0; i < count; i++ {
			substr[i] = c
		}
		heap.Push(h, string(substr))
	}
	return h.Stringify()
}
