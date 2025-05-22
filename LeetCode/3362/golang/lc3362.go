package lc3362

import "container/heap"

type Item struct {
	value    int
	priority int
	index    int
}
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(lhs, rhs int) bool {
	return pq[lhs].priority > pq[rhs].priority
}

func (pq PriorityQueue) Swap(lhs, rhs int) {
	pq[lhs], pq[rhs] = pq[rhs], pq[lhs]
	pq[lhs].index = lhs
	pq[rhs].index = rhs
}

func (pq *PriorityQueue) Push(x any) {
	length := len(*pq)
	item := x.(*Item)
	item.index = length
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	instance := *pq
	length := len(instance)
	item := instance[length-1]
	instance[length-1] = nil
	item.index = -1
	*pq = instance[0 : length-1]
	return item
}

func (pq *PriorityQueue) update(item *Item, value int, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

// Runtime complexity:
// Auxiliary space:
func maxRemoval(nums []int, queries [][]int) int {
	ans := 0
	return ans
}
