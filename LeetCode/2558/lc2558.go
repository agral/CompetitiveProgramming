package lc2558

import (
	"container/heap"
	"math"
)

type Item struct {
	value    int
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(a, b int) bool {
	// Pop() should return the biggest - not lowest - priority;
	// so greater-than operator is used.
	return pq[a].priority > pq[b].priority
}

func (pq PriorityQueue) Swap(a, b int) {
	pq[a], pq[b] = pq[b], pq[a]
	pq[a].index = a
	pq[b].index = b
}

func (pq *PriorityQueue) Push(x any) {
	l := len(*pq)
	item := x.(*Item)
	item.index = l
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	l := len(old)
	item := old[l-1]
	item.index = -1 // just for safety
	*pq = old[0 : l-1]
	return item
}

// update() modifies the value and the priority of an Item already in the queue.
func (pq *PriorityQueue) update(item *Item, value int, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func pickGifts(gifts []int, k int) int64 {
	pq := make(PriorityQueue, 0)
	for _, pile := range gifts {
		heap.Push(&pq, &Item{value: 0, priority: pile})
	}

	for i := 0; i < k; i++ {
		pile := heap.Pop(&pq).(*Item)
		item := &Item{value: 0, priority: int(math.Sqrt(float64(pile.priority)))}
		heap.Push(&pq, item)
	}

	ans := int64(0)
	l := pq.Len()
	for i := 0; i < l; i++ {
		item := heap.Pop(&pq).(*Item)
		ans += int64(item.priority)
	}
	return ans
}
