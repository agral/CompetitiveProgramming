package lc3446

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

func Test_sortMatrix(t *testing.T) {
	testcases := []struct {
		grid     [][]int
		expected [][]int
	}{
		{[][]int{{1, 7, 3}, {9, 8, 2}, {4, 5, 6}}, [][]int{{8, 2, 3}, {9, 6, 7}, {4, 5, 1}}},
		{[][]int{{0, 1}, {1, 2}}, [][]int{{2, 1}, {1, 0}}},
	}

	for i, tc := range testcases {
		actual := sortMatrix(tc.grid)
		if !IsEqual(actual, tc.expected) {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.grid, tc.expected, actual)
		}
	}
}
