package lc2593

import (
	"container/heap"
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
	// need to resolve ties so that elements with lower indices get out first.
	if pq[a].priority == pq[b].priority {
		return pq[a].value < pq[b].value
	}
	// otherwise it's a standard heap (priority queue with Less-Than prority)
	return pq[a].priority < pq[b].priority
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

func findScore(nums []int) int64 {
	ans := int64(0)
	pq := make(PriorityQueue, 0)
	seen := make([]bool, len(nums))
	for idx, val := range nums {
		heap.Push(&pq, &Item{value: idx, priority: val})
		seen[idx] = false
	}

	for i := 0; i < len(nums); i++ {
		smallest := heap.Pop(&pq).(*Item)
		if !seen[smallest.value] {
			ans += int64(smallest.priority)
			seen[smallest.value] = true
			if smallest.value > 0 {
				seen[smallest.value-1] = true
			}
			if smallest.value < len(nums)-1 {
				seen[smallest.value+1] = true
			}
		}
	}

	return ans
}
