package lc0883

import "testing"

func Test_projectionArea(t *testing.T) {
	testcases := []struct {
		grid     [][]int
		expected int
	}{
		{[][]int{{1, 2}, {3, 4}}, 17}, // 4 + 6 + 7
		{[][]int{{2}}, 5},             // 1 + 2 + 2
		{[][]int{{1, 0}, {0, 2}}, 8},
	}

	for i, tc := range testcases {
		actual := projectionArea(tc.grid)
		if actual != tc.expected {
			t.Errorf("Testcase projectionArea#%02d (%v) failed: want %d, got %d",
				i+1, tc.grid, tc.expected, actual)
		}
	}
}
