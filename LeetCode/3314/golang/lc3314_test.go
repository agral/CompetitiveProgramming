package lc3314

import (
	"slices"
	"testing"
)

func Test_minBitwiseArray(t *testing.T) {
	testcases := []struct {
		nums     []int
		expected []int
	}{
		{[]int{2, 3, 5, 7}, []int{-1, 1, 4, 3}},
		{[]int{11, 13, 31}, []int{9, 12, 15}},
	}

	for i, tc := range testcases {
		actual := minBitwiseArray(tc.nums)
		if !slices.Equal(actual, tc.expected) {
			t.Errorf("Testcase %02d (%v) failed: want %v, got %v", i+1, tc.nums, tc.expected, actual)
		}
	}
}
