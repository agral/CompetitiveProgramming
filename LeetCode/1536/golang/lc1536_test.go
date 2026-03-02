package lc1536

import "testing"

func Test_minSwaps(t *testing.T) {
	testcases := []struct {
		grid     [][]int
		expected int
	}{
		{[][]int{ // example 1
			{0, 0, 1},
			{1, 1, 0},
			{1, 0, 0},
		}, 3},
		{[][]int{ // example 2
			{0, 1, 1, 0},
			{0, 1, 1, 0},
			{0, 1, 1, 0},
			{0, 1, 1, 0},
		}, -1},
		{[][]int{ //example 3
			{1, 0, 0},
			{1, 1, 0},
			{1, 1, 1},
		}, 0},
	}

	for i, tc := range testcases {
		actual := minSwaps(tc.grid)
		if actual != tc.expected {
			t.Errorf("Testcase minSwaps#%02d (%v) failed: want %d, got %d",
				i+1, tc.grid, tc.expected, actual)
		}
	}
}
