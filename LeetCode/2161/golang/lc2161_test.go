package lc2161

import (
	"slices"
	"testing"
)

func Test_pivotArray(t *testing.T) {
	testcases := []struct {
		nums     []int
		pivot    int
		expected []int
	}{
		{[]int{9, 12, 5, 10, 14, 3, 10}, 10, []int{9, 5, 3, 10, 10, 12, 14}},
		{[]int{-3, 4, 3, 2}, 2, []int{-3, 2, 4, 3}},
	}

	for i, tc := range testcases {
		actual := pivotArray(tc.nums, tc.pivot)
		if !slices.Equal(actual, tc.expected) {
			t.Errorf("Testcase %02d (%v|%d) failed: want %v, got %v",
				i+1, tc.nums, tc.pivot, tc.expected, actual)
		}
	}
}
