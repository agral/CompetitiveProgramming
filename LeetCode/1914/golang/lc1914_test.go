package lc1914

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

func Test_rotateGrid(t *testing.T) {
	testcases := []struct {
		grid     [][]int
		k        int
		expected [][]int
	}{
		{
			[][]int{
				{40, 10},
				{30, 20},
			}, 1,
			[][]int{
				{10, 20},
				{40, 30},
			},
		},
		{
			[][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
				{13, 14, 15, 16},
			}, 2,
			[][]int{
				{3, 4, 8, 12},
				{2, 11, 10, 16},
				{1, 7, 6, 15},
				{5, 9, 13, 14},
			},
		},
		{
			[][]int{
				{1, 2, 3, 4},
				{16, 1, 2, 5},
				{15, 8, 3, 6},
				{14, 7, 4, 7},
				{13, 6, 5, 8},
				{12, 11, 10, 9},
			}, 1,
			[][]int{
				{2, 3, 4, 5},
				{1, 2, 3, 6},
				{16, 1, 4, 7},
				{15, 8, 5, 8},
				{14, 7, 6, 9},
				{13, 12, 11, 10},
			},
		},
	}

	for i, tc := range testcases {
		actual := rotateGrid(tc.grid, tc.k)
		if !IsEqual(actual, tc.expected) {
			t.Errorf("Testcase rotateGrid#%02d (%v) failed: want %v, got %v",
				i+1, tc.grid, tc.expected, actual)
		}
	}
}
