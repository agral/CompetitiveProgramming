package lc1266

import "testing"

func Test_minTimeToVisitAllPoints(t *testing.T) {
	testcases := []struct {
		points   [][]int
		expected int
	}{
		{[][]int{{1, 1}, {3, 4}, {-1, 0}}, 7},
		{[][]int{{3, 2}, {-2, 2}}, 5},
	}

	for i, tc := range testcases {
		actual := minTimeToVisitAllPoints(tc.points)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.points, tc.expected, actual)
		}
	}
}
