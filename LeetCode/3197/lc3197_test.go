package lc3197

import "testing"

func Test_minimumSum(t *testing.T) {
	testcases := []struct {
		grid     [][]int
		expected int
	}{
		{[][]int{{1, 0, 1}, {1, 1, 1}}, 5},
		{[][]int{{1, 0, 1, 0}, {0, 1, 0, 1}}, 5},
	}

	for i, tc := range testcases {
		actual := minimumSum(tc.grid)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.grid, tc.expected, actual)
		}
	}
}
