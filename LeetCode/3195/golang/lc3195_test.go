package lc3195

import "testing"

func Test_minimumArea(t *testing.T) {
	testcases := []struct {
		grid     [][]int
		expected int
	}{
		{[][]int{{0, 1, 0}, {1, 0, 1}}, 6},
		{[][]int{{1, 0}, {0, 0}}, 1},
	}

	for i, tc := range testcases {
		actual := minimumArea(tc.grid)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.grid, tc.expected, actual)
		}
	}
}
