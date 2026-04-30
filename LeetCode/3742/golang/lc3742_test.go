package lc3742

import "testing"

func Test_maxPathScore(t *testing.T) {
	testcases := []struct {
		grid     [][]int
		k        int
		expected int
	}{
		{[][]int{{0, 1}, {2, 0}}, 1, 2},
		{[][]int{{0, 1}, {1, 2}}, 1, -1},
	}

	for i, tc := range testcases {
		actual := maxPathScore(tc.grid, tc.k)
		if actual != tc.expected {
			t.Errorf("Testcase maxPathScore#%02d (%v | %d) failed: want %d, got %d",
				i+1, tc.grid, tc.k, tc.expected, actual)
		}
	}
}
