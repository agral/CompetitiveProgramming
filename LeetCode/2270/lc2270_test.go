package lc2270

import (
	"testing"
)

func TestWaysToSplitArray(t *testing.T) {
	testingData := []struct {
		nums     []int
		expected int
	}{
		{nums: []int{10, 4, -8, 7}, expected: 2},
		{nums: []int{2, 3, 1, 0}, expected: 2},
	}

	for i, d := range testingData {
		actual := waysToSplitArray(d.nums)
		if actual != d.expected {
			t.Errorf("Testcase #%02d (%v): want %d, got %d", i, d.nums, d.expected, actual)
		}
	}
}
