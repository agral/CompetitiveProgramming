package lc3459

import "testing"

func Test_lenOfVDiagonal(t *testing.T) {
	testcases := []struct {
		grid     [][]int
		expected int
	}{
		{[][]int{
			{2, 2, 1, 2, 2},
			{2, 0, 2, 2, 0},
			{2, 0, 1, 1, 0},
			{1, 0, 2, 2, 2},
			{2, 0, 0, 2, 2},
		}, 5},

		{[][]int{
			{2, 2, 2, 2, 2},
			{2, 0, 2, 2, 0},
			{2, 0, 1, 1, 0},
			{1, 0, 2, 2, 2},
			{2, 0, 0, 2, 2},
		}, 4},

		{[][]int{
			{1, 2, 2, 2, 2},
			{2, 2, 2, 2, 0},
			{2, 0, 0, 0, 0},
			{0, 0, 2, 2, 2},
			{2, 0, 0, 2, 0},
		}, 5},
	}

	for i, tc := range testcases {
		actual := lenOfVDiagonal(tc.grid)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.grid, tc.expected, actual)
		}
	}
}
