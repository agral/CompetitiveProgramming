package lc0066

import (
	"slices"
	"testing"
)

func Test_plusOne(t *testing.T) {
	testcases := []struct {
		digits   []int
		expected []int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 4}},
		{[]int{4, 3, 2, 1}, []int{4, 3, 2, 2}},
		{[]int{9}, []int{1, 0}},
		{[]int{0}, []int{1}},
		{[]int{1, 9}, []int{2, 0}},
		{[]int{8, 9}, []int{9, 0}},
		{[]int{9, 9}, []int{1, 0, 0}},
	}

	for i, tc := range testcases {
		actual := plusOne(tc.digits)
		if !slices.Equal(actual, tc.expected) {
			t.Errorf("Testcase %02d (%v) failed: want %v, got %v",
				i+1, tc.digits, tc.expected, actual)
		}
	}
}
