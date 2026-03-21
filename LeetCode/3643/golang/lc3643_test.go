package lc3643

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

func Test_reverseSubmatrix(t *testing.T) {
	testcases := []struct {
		grid     [][]int
		x        int
		y        int
		k        int
		expected [][]int
	}{
		{
			[][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
				{13, 14, 15, 16},
			},
			1, 0, 3,
			[][]int{
				{1, 2, 3, 4},
				{13, 14, 15, 8},
				{9, 10, 11, 12},
				{5, 6, 7, 16},
			},
		},
		{
			[][]int{
				{3, 4, 2, 3},
				{2, 3, 4, 2},
			},
			0, 2, 2,
			[][]int{
				{3, 4, 4, 2},
				{2, 3, 2, 3},
			},
		},
	}

	for i, tc := range testcases {
		actual := reverseSubmatrix(tc.grid, tc.x, tc.y, tc.k)
		if !IsEqual(actual, tc.expected) {
			t.Errorf("Testcase reverseSubmatrix#%02d (%v) failed.\nwant:\n%v\ngot:\n%v",
				i+1, tc.grid, tc.expected, actual)
		}
	}
}
