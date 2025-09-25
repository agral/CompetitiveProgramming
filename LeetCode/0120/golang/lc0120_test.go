package lc0120

import "testing"

func Test_minimumTotal(t *testing.T) {
	testcases := []struct {
		triangle [][]int
		expected int
	}{
		{[][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}, 11},
		{[][]int{{-10}}, -10},
		{[][]int{{-1}, {2, 3}, {1, -1, -3}}, -1},
	}

	for i, tc := range testcases {
		actual := minimumTotal(tc.triangle)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.triangle, tc.expected, actual)
		}
	}
}
