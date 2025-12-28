package lc1351

import "testing"

func Test_countNegatives(t *testing.T) {
	testcases := []struct {
		grid     [][]int
		expected int
	}{
		{[][]int{
			{4, 3, 2, -1},
			{3, 2, 1, -1},
			{1, 1, -1, -2},
		}, 4},
		{[][]int{ // all positives:
			{3, 2},
			{1, 0},
		}, 0},
		{[][]int{ // all negatives:
			{-1, -2, -3},
			{-3, -4, -5},
		}, 6},
	}

	for i, tc := range testcases {
		actual := countNegatives(tc.grid)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.grid, tc.expected, actual)
		}
	}
}
