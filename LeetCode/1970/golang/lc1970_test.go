package lc1970

import "testing"

func Test_latestDayToCross(t *testing.T) {
	testcases := []struct {
		row      int
		col      int
		cells    [][]int
		expected int
	}{
		{2, 2, [][]int{{1, 1}, {2, 1}, {1, 2}, {2, 2}}, 2},
		{2, 2, [][]int{{1, 1}, {1, 2}, {2, 1}, {2, 2}}, 1},
		{3, 3, [][]int{{1, 2}, {2, 1}, {3, 3}, {2, 2}, {1, 1}, {1, 3}, {2, 3}, {3, 2}, {3, 1}}, 3},
	}

	for i, tc := range testcases {
		actual := latestDayToCross(tc.row, tc.col, tc.cells)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%dx%d, %v) failed: want %d, got %d",
				i+1, tc.row, tc.col, tc.cells, tc.expected, actual)
		}
	}
}
