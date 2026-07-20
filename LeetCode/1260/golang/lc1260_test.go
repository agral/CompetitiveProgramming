package lc1260

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

func Test_shiftGrid(t *testing.T) {
	testcases := []struct {
		grid     [][]int
		k        int
		expected [][]int
	}{
		{
			[][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			1,
			[][]int{
				{9, 1, 2},
				{3, 4, 5},
				{6, 7, 8},
			},
		},
	}

	for i, tc := range testcases {
		actual := shiftGrid(tc.grid, tc.k)
		if !IsEqual(actual, tc.expected) {
			t.Errorf("Testcase shiftGrid#%02d (%v, %d) failed: want %v, got %v",
				i+1, tc.grid, tc.k, tc.expected, actual)
		}
	}
}
