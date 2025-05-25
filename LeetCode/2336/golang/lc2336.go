package lc2336

import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	instance := *h
	n := len(instance)
	ans := instance[n-1]
	*h = instance[0 : n-1]
	return ans
}

type SmallestInfiniteSet struct {
	nextNum int
	added   IntHeap
}

func Constructor() SmallestInfiniteSet {
	h := &IntHeap{}
	heap.Init(h)
	instance := SmallestInfiniteSet{
		nextNum: 1,
		added:   *h,
	}
	return instance
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	if this.added.Len() == 0 {
		ans := this.nextNum
		this.nextNum += 1
		return ans
	}
	ans := heap.Pop(&this.added).(int)
	return ans
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	if num < this.nextNum {
		heap.Push(&this.added, num)
	}
}
