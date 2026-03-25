package lc3546

import "testing"

func Test_canPartitionGrid(t *testing.T) {
	testcases := []struct {
		grid     [][]int
		expected bool
	}{
		{
			[][]int{
				{1, 4},
				{2, 3},
			}, true,
		},
		{
			[][]int{
				{1, 3},
				{2, 4},
			}, false,
		},
		{
			[][]int{
				{1, 1, 1, 1, 1, 5},
				{1, 1, 1, 1, 1, 5},
				{1, 1, 1, 1, 1, 5},
			}, true,
		},
		{
			[][]int{
				{5, 1, 1, 1, 1, 1},
				{5, 1, 1, 1, 1, 1},
				{5, 1, 1, 1, 1, 1},
			}, true,
		},
		{
			[][]int{
				{5, 1, 1, 1, 1, 1},
				{5, 1, 1, 1, 1, 1},
				{5, 2, 1, 1, 1, 1},
			}, false,
		},
	}

	for i, tc := range testcases {
		actual := canPartitionGrid(tc.grid)
		if actual != tc.expected {
			t.Errorf("Testcase canPartitionGrid#%02d (%v) failed: want %t, got %t",
				i+1, tc.grid, tc.expected, actual)
		}
	}
}
