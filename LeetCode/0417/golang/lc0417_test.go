package lc0417

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

func Test_pacificAtlantic(t *testing.T) {
	testcases := []struct {
		heights  [][]int
		expected [][]int
	}{
		{
			[][]int{
				{1, 2, 2, 3, 5},
				{3, 2, 3, 4, 4},
				{2, 4, 5, 3, 1},
				{6, 7, 1, 4, 5},
				{5, 1, 1, 2, 4},
			}, [][]int{
				{0, 4}, {1, 3}, {1, 4}, {2, 2}, {3, 0}, {3, 1}, {4, 0},
			},
		},
		{
			[][]int{
				{1},
			}, [][]int{
				{0, 0},
			},
		},
	}

	for i, tc := range testcases {
		actual := pacificAtlantic(tc.heights)
		if !IsEqual(tc.expected, actual) {
			t.Errorf("Testcase %02d failed: want %v, got %v", i+1, tc.expected, actual)
		}
	}
}
