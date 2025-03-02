package lc2570

import (
	"slices"
	"testing"
)

func isEqual(lhs [][]int, rhs [][]int) bool {
	SZ := len(lhs)
	if SZ != len(rhs) {
		return false
	}

	for i := range SZ {
		if !slices.Equal(lhs[i], rhs[i]) {
			return false
		}
	}
	return true
}

func Test_mergeArrays(t *testing.T) {
	testcases := []struct {
		nums1    [][]int
		nums2    [][]int
		expected [][]int
	}{
		{
			[][]int{{1, 2}, {2, 3}, {4, 5}},
			[][]int{{1, 4}, {3, 2}, {4, 1}},
			[][]int{{1, 6}, {2, 3}, {3, 2}, {4, 6}},
		},
		{
			[][]int{{2, 4}, {3, 6}, {5, 5}},
			[][]int{{1, 3}, {4, 3}},
			[][]int{{1, 3}, {2, 4}, {3, 6}, {4, 3}, {5, 5}},
		},
	}

	for i, tc := range testcases {
		actual := mergeArrays(tc.nums1, tc.nums2)
		if !isEqual(actual, tc.expected) {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.nums1, tc.expected, actual)
		}
	}
}
