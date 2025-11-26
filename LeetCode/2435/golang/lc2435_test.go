package lc2435

import "testing"

func Test_numberOfPaths(t *testing.T) {
	testcases := []struct {
		grid     [][]int
		k        int
		expected int
	}{
		{[][]int{{5, 2, 4}, {3, 0, 5}, {0, 7, 2}}, 3, 2},
		{[][]int{{0, 0}}, 5, 1},
		{[][]int{{7, 3, 4, 9}, {2, 3, 6, 2}, {2, 3, 7, 0}}, 1, 10},
	}

	for i, tc := range testcases {
		actual := numberOfPaths(tc.grid, tc.k)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v | %d) failed: want %d, got %d",
				i+1, tc.grid, tc.k, tc.expected, actual)
		}
	}
}
