package lc1267

import "testing"

func TestCountServers(t *testing.T) {
	testcases := []struct {
		grid     [][]int
		expected int
	}{
		{[][]int{{1, 0}, {0, 1}}, 0},
		{[][]int{{0, 1}, {1, 0}}, 0},
		{[][]int{{1, 0}, {1, 1}}, 3},
		{[][]int{{1, 1, 0, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}, {0, 0, 0, 1}}, 4},
		{[][]int{{1, 0, 1}}, 2},
	}

	for i, tc := range testcases {
		actual := countServers(tc.grid)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.grid, tc.expected, actual)
		}
	}
}
