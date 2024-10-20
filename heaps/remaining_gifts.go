package heaps

import (
	"container/heap"
	"math"
)

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func giftsTotal(h *IntHeap) int {
	total := 0
	for _, g := range *h {
		total += -1 * g
	}
	return total
}

func remainingGifts(gifts []int, k int) int {
	if len(gifts) < 1 {
		return 0
	}
	h := &IntHeap{}
	heap.Init(h)

	for _, g := range gifts {
		heap.Push(h, -g)
	}
	if k == 0 {
		return giftsTotal(h)
	}

	// the last item will be the largest
	for i := 0; i < k; i++ {
		mostGifts := -1 * heap.Pop(h).(int)
		if mostGifts > 0 {
			remainingGifts := math.Sqrt(float64(mostGifts))
			heap.Push(h, -1*int(remainingGifts))
		}
	}

	return giftsTotal(h)
}
