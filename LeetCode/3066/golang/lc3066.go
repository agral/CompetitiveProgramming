package lc3066

import "container/heap"

type IntHeap []int // a min-heap of int values

func (h IntHeap) Len() int {
	return len(h)
}
func (h IntHeap) Less(a, b int) bool {
	return h[a] < h[b]
}
func (h IntHeap) Swap(a, b int) {
	h[a], h[b] = h[b], h[a]
}
func (h *IntHeap) Push(val any) {
	*h = append(*h, val.(int))
}
func (h *IntHeap) Pop() any {
	old := *h
	l := len(old)
	ans := old[l-1]
	*h = old[:l-1]
	return ans
}

func minOperations(nums []int, k int) int {
	minHeap := &IntHeap{}
	heap.Init(minHeap)
	for _, num := range nums {
		heap.Push(minHeap, num)
	}
	ans := 0

	for minHeap.Len() > 1 && (*minHeap)[0] < k {
		x := heap.Pop(minHeap).(int)
		y := heap.Pop(minHeap).(int)
		// Insert (min(x, y) * 2 + max(x, y)). Guaranteed that min(x, y) == x, max(x, y) == y.
		heap.Push(minHeap, 2*x+y)
		ans++
	}

	return ans
}
