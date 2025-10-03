package lc0407

import "testing"

func Test_trapRainWater(t *testing.T) {
	testcases := []struct {
		heightMap [][]int
		expected  int
	}{
		{[][]int{
			{1, 4, 3, 1, 3, 2},
			{3, 2, 1, 3, 2, 4},
			{2, 3, 3, 2, 3, 1},
		}, 4},

		{[][]int{
			{3, 3, 3, 3, 3},
			{3, 2, 2, 2, 3},
			{3, 2, 1, 2, 3},
			{3, 2, 2, 2, 3},
			{3, 3, 3, 3, 3},
		}, 10},

		{[][]int{
			{12, 13, 1, 12},
			{13, 4, 13, 12},
			{13, 8, 10, 12},
			{12, 13, 12, 12},
			{13, 13, 13, 13},
		}, 14},
	}

	for i, tc := range testcases {
		actual := trapRainWater(tc.heightMap)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.heightMap, tc.expected, actual)
		}
	}
}
