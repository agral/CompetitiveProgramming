package lc1594

import "testing"

func Test_maxProductPath(t *testing.T) {
	testcases := []struct {
		grid     [][]int
		expected int
	}{
		{
			[][]int{
				{-1, -2, -3},
				{-2, -3, -3},
				{-3, -3, -2},
			}, -1,
		},
		{
			[][]int{
				{1, -2, 1},
				{1, -2, 1},
				{3, -4, 1},
			}, 8,
		},
		{
			[][]int{
				{1, 3},
				{0, -4},
			}, 0,
		},
	}

	for i, tc := range testcases {
		actual := maxProductPath(tc.grid)
		if actual != tc.expected {
			t.Errorf("Testcase maxProductPath#%02d (%v) failed: want %d, got %d",
				i+1, tc.grid, tc.expected, actual)
		}
	}
}
