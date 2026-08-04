package lc3731

import (
	"slices"
	"testing"
)

func Test_findMissingElements(t *testing.T) {
	testcases := []struct {
		nums     []int
		expected []int
	}{
		{[]int{1, 4, 2, 5}, []int{3}},
		{[]int{7, 8, 6, 9}, []int{}},
		{[]int{5, 1}, []int{2, 3, 4}},
	}

	for i, tc := range testcases {
		actual := findMissingElements(tc.nums)
		if !slices.Equal(actual, tc.expected) {
			t.Errorf("Testcase findMissingElements#%02d (%v) failed: want %v, got %v",
				i+1, tc.nums, tc.expected, actual)
		}
	}
}
