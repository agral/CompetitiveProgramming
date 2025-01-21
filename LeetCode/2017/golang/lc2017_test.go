package lc2017

import "testing"

func TestGridGame(t *testing.T) {
	testcases := []struct {
		grid     [][]int
		expected int64
	}{
		{[][]int{{2, 5, 4}, {1, 5, 1}}, 4},
		{[][]int{{3, 3, 1}, {8, 5, 2}}, 4},
		{[][]int{{1, 3, 1, 15}, {1, 3, 3, 1}}, 7},
	}

	for i, tc := range testcases {
		actual := gridGame(tc.grid)
		if actual != tc.expected {
			t.Errorf("Testcase %02d failed: want %d, got %d", i+1, tc.expected, actual)
		}
	}
}
