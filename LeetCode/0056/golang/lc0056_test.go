package lc0056

import (
	"slices"
	"testing"
)

func IsEqual(lhs [][]int, rhs [][]int) bool {
	if len(lhs) != len(rhs) {
		return false
	}
	for i := range lhs {
		if !slices.Equal(lhs[i], rhs[i]) {
			return false
		}
	}
	return true
}

func Test_merge(t *testing.T) {
	testcases := []struct {
		intervals [][]int
		expected  [][]int
	}{
		{
			[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			[][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			[][]int{{1, 4}, {4, 5}},
			[][]int{{1, 5}},
		},
		{
			[][]int{{4, 7}, {1, 4}},
			[][]int{{1, 7}},
		},
		{
			[][]int{{2, 3}, {2, 2}, {3, 3}, {1, 3}, {5, 7}, {2, 2}, {4, 6}},
			[][]int{{1, 3}, {4, 7}},
		},
	}

	for i, tc := range testcases {
		actual := merge(tc.intervals)
		if !IsEqual(actual, tc.expected) {
			t.Errorf("Testcase merge#%02d (%v) failed: want %v, got %v",
				i+1, tc.intervals, tc.expected, actual)
		}
	}
}
