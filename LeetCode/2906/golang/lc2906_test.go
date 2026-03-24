package lc2906

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

func Test_constructProductMatrix(t *testing.T) {
	testcases := []struct {
		grid     [][]int
		expected [][]int
	}{
		{
			[][]int{
				{1, 2},
				{3, 4},
			},
			[][]int{
				{24, 12},
				{8, 6},
			},
		},
		{
			[][]int{
				{12345, 2, 1},
			},
			[][]int{
				{2, 0, 0},
			},
		},
	}

	for i, tc := range testcases {
		actual := constructProductMatrix(tc.grid)
		if !IsEqual(actual, tc.expected) {
			t.Errorf("Testcase constructProductMatrix#%02d (%v) failed.\nwant:\n%v\ngot:\n%v",
				i+1, tc.grid, tc.expected, actual)
		}
	}
}
