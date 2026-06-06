package lc2574

import (
	"slices"
	"testing"
)

func Test_leftRightDifference(t *testing.T) {
	testcases := []struct {
		nums     []int
		expected []int
	}{
		{
			[]int{10, 4, 8, 3},
			[]int{15, 1, 11, 22},
		},
		{
			[]int{1},
			[]int{0},
		},
	}

	for i, tc := range testcases {
		actual := leftRightDifference(tc.nums)
		if !slices.Equal(actual, tc.expected) {
			t.Errorf("Testcase leftRightDifference#%02d (%v) failed: want %v, got %v",
				i+1, tc.nums, tc.expected, actual)
		}
	}
}
