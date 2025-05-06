package lc1920

import (
	"slices"
	"testing"
)

func Test_buildArray(t *testing.T) {
	testcases := []struct {
		nums     []int
		expected []int
	}{
		{[]int{0, 2, 1, 5, 3, 4}, []int{0, 1, 2, 4, 5, 3}},
		{[]int{5, 0, 1, 2, 3, 4}, []int{4, 5, 0, 1, 2, 3}},
	}

	for i, tc := range testcases {
		actual := buildArray(tc.nums)
		if !slices.Equal(actual, tc.expected) {
			t.Errorf("Testcase %02d (%v) failed: want %v, got %v", i+1, tc.nums, tc.expected, actual)
		}
	}
}
