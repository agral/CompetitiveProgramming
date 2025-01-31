package lc0827

import "testing"

func Test_largestIsland(t *testing.T) {
	testcases := []struct {
		grid     [][]int
		expected int
	}{
		{[][]int{{1, 0}, {0, 1}}, 3},
		{[][]int{{1, 1}, {1, 0}}, 4},
		{[][]int{{1, 1}, {1, 1}}, 4},
	}

	for i, tc := range testcases {
		actual := largestIsland(tc.grid)
		if actual != tc.expected {
			t.Errorf("Testcase %02d failed: want %d, got %d", i+1, tc.expected, actual)
		}
	}
}
