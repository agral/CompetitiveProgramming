package lc3379

import (
	"slices"
	"testing"
)

func Test_constructTransformedArray(t *testing.T) {
	testcases := []struct {
		nums     []int
		expected []int
	}{
		{[]int{3, -2, 1, 1}, []int{1, 1, 1, 3}},
		{[]int{-1, 4, -1}, []int{-1, -1, 4}},
	}

	for i, tc := range testcases {
		actual := constructTransformedArray(tc.nums)
		if !slices.Equal(actual, tc.expected) {
			t.Errorf("Testcase constructTransformedArray#%02d (%v) failed: want %v, got %v",
				i+1, tc.nums, tc.expected, actual)
		}
	}
}
