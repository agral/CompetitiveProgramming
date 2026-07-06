package lc1288

import "testing"

func Test_removeCoveredIntervals(t *testing.T) {
	testcases := []struct {
		intervals [][]int
		expected  int
	}{
		{[][]int{{1, 4}, {3, 6}, {2, 8}}, 2},         // interval {3, 6} is removed.
		{[][]int{{1, 4}, {2, 3}}, 1},                 // interval {2, 3} is removed.
		{[][]int{{1, 2}, {1, 3}, {1, 4}, {1, 5}}, 1}, // only the {1, 5} interval is unique.
		{[][]int{{4, 5}, {3, 5}, {2, 5}, {1, 5}}, 1}, // again, only the {1, 5} is unique.

		// {1, 4} and {4, 7} cover the entire range, but there's no overlap, so all have to be used.
		{[][]int{{1, 4}, {2, 5}, {3, 6}, {4, 7}}, 4},
	}

	for i, tc := range testcases {
		actual := removeCoveredIntervals(tc.intervals)
		if actual != tc.expected {
			t.Errorf("Testcase removeCoveredIntervals#%02d (%v) failed: want %d, got %d",
				i+1, tc.intervals, tc.expected, actual)
		}
	}
}
