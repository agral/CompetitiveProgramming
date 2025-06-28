package lc2099

import (
	"slices"
	"testing"
)

func Test_maxSubsequence(t *testing.T) {
	testcases := []struct {
		nums     []int
		k        int
		expected []int
	}{
		{[]int{2, 1, 3, 3}, 2, []int{3, 3}},
		{[]int{-1, -2, 3, 4}, 3, []int{-1, 3, 4}},
		{[]int{3, 4, 3, 3}, 2, []int{3, 4}},
	}

	for i, tc := range testcases {
		actual := maxSubsequence(tc.nums, tc.k)
		if !slices.Equal(actual, tc.expected) {
			t.Errorf("Testcase %02d (%v | %d) failed: want %v, got %v",
				i+1, tc.nums, tc.k, tc.expected, actual)
		}
	}
}
