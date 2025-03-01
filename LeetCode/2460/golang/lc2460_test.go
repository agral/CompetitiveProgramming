package lc2460

import (
	"slices"
	"testing"
)

func Test_applyOperations(t *testing.T) {
	testcases := []struct {
		nums     []int
		expected []int
	}{
		{[]int{1, 2, 2, 1, 1, 0}, []int{1, 4, 2, 0, 0, 0}},
		{[]int{0, 1}, []int{1, 0}}, // no operations are performed, zero is shifted to the end.
		// try some more interesting test cases:
		{[]int{0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0}}, // only zeroes!
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}}, // no zeroes, no operations can be made.
		{[]int{1, 1, 1, 1}, []int{2, 2, 0, 0}},       // no zeroes initially, but all numbers match.
	}

	for i, tc := range testcases {
		actual := applyOperations(tc.nums)
		if !slices.Equal(actual, tc.expected) {
			t.Errorf("Testcase %02d (%v) failed: want %v, got %v", i+1, tc.nums, tc.expected, actual)
		}
	}
}
