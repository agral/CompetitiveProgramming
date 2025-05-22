package lc0073

import (
	"slices"
	"testing"
)

func areEqual(lhs [][]int, rhs [][]int) bool {
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

func Test_setZeroes(t *testing.T) {
	testcases := []struct {
		matrix   [][]int
		expected [][]int
	}{
		{[][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}, [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}},
		{[][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}, [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}}},
	}

	for i, tc := range testcases {
		setZeroes(tc.matrix)
		if !areEqual(tc.matrix, tc.expected) {
			t.Errorf("Testcase %02d failed: want %d, got %d", i+1, tc.expected, tc.matrix)
		}
	}
}
