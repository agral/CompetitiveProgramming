package lc2411

import (
	"slices"
	"testing"
)

func Test_smallestSubarrays(t *testing.T) {
	testcases := []struct {
		nums     []int
		expected []int
	}{
		{[]int{1, 0, 2, 1, 3}, []int{3, 3, 2, 2, 1}},
		{[]int{1, 2}, []int{2, 1}},
	}

	for i, tc := range testcases {
		actual := smallestSubarrays(tc.nums)
		if !slices.Equal(actual, tc.expected) {
			t.Errorf("Testcase %02d (%v) failed: want %v, got %v", i+1, tc.nums, tc.expected, actual)
		}
	}
}
