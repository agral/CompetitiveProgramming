package lc0162

import (
	"slices"
	"testing"
)

func Test_findPeakElement(t *testing.T) {
	testcases := []struct {
		nums     []int
		expected []int // findPeakElement should return a number that's in expected.
	}{
		{[]int{1, 2, 3, 1}, []int{2}},
		{[]int{1, 2, 1, 3, 5, 6, 4}, []int{1, 5}},
		{[]int{0, 1, 2, 3, 4, 5, 6}, []int{6}},
		{[]int{6, 5, 4, 3, 2, 1, 0}, []int{0}},
	}

	for i, tc := range testcases {
		actual := findPeakElement(tc.nums)
		if !slices.Contains(tc.expected, actual) {
			t.Errorf("Testcase %02d (%v) failed: accepting any of %v, got %d", i+1, tc.nums, tc.expected, actual)
		}
	}
}
