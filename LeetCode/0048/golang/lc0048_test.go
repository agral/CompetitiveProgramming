package lc0048

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

func Test_rotate(t *testing.T) {
	testcases := []struct {
		matrix   [][]int
		expected [][]int
	}{
		{
			[][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9}},
			[][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3}},
		},
		{
			[][]int{
				{5, 1, 9, 11},
				{2, 4, 8, 10},
				{13, 3, 6, 7},
				{15, 14, 12, 16}},
			[][]int{
				{15, 13, 2, 5},
				{14, 3, 4, 1},
				{12, 6, 8, 9},
				{16, 7, 10, 11}},
		},
	}

	for i, tc := range testcases {
		rotate(tc.matrix)
		if !IsEqual(tc.matrix, tc.expected) {
			t.Errorf("Testcase rotate#%02d failed: want %d, got %d",
				i+1, tc.matrix, tc.expected)
		}
	}
}
