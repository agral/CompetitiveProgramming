package lc0118

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

func Test_generate(t *testing.T) {
	testcases := []struct {
		numRows  int
		expected [][]int
	}{
		{5, [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}},
		{1, [][]int{{1}}},
	}

	for i, tc := range testcases {
		actual := generate(tc.numRows)
		if !IsEqual(actual, tc.expected) {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.numRows, tc.expected, actual)
		}
	}
}
