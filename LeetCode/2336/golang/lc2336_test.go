package lc2336

import (
	"container/heap"
	"testing"
)

func Test_SmallestInfiniteSet(t *testing.T) {

	assertEqual := func(actual, expected int) {
		if actual != expected {
			t.Errorf("Assertion failed. Expected %d, got %d", expected, actual)
		}
	}
	s := Constructor()
	var received int
	s.AddBack(2) // 2 is already in the set, so this is a no-op.
	received = s.PopSmallest()
	assertEqual(received, 1)
	received = s.PopSmallest()
	assertEqual(received, 2)
	received = s.PopSmallest()
	assertEqual(received, 3)
	s.AddBack(1)
	received = s.PopSmallest()
	assertEqual(received, 1)
	received = s.PopSmallest()
	assertEqual(received, 4)
	received = s.PopSmallest()
	assertEqual(received, 5)
}

func Test_SmallestInfiniteSet2(t *testing.T) {
	assertEqual := func(actual, expected int) {
		if actual != expected {
			t.Errorf("Assertion failed. Expected %d, got %d", expected, actual)
		} else {
			//t.Logf("Assertion %d == %d passed", expected, actual)
		}
	}
	s := Constructor()
	var received int
	received = s.PopSmallest()
	assertEqual(received, 1)
	s.AddBack(1)
	received = s.PopSmallest()
	assertEqual(received, 1)
	received = s.PopSmallest()
	assertEqual(received, 2)
	received = s.PopSmallest()
	assertEqual(received, 3)
	s.AddBack(2)
	s.AddBack(3)
	received = s.PopSmallest()
	assertEqual(received, 2)
	received = s.PopSmallest()
	assertEqual(received, 3)
}

func Test_IntHeap(t *testing.T) {
	assertEqual := func(actual, expected int) {
		if actual != expected {
			t.Errorf("Assertion failed. Expected %d, got %d", expected, actual)
		}
	}
	h := &IntHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)
	assertEqual(heap.Pop(h).(int), 1)
	assertEqual(heap.Pop(h).(int), 2)
	assertEqual(heap.Pop(h).(int), 3)
	assertEqual(heap.Pop(h).(int), 5)
}
