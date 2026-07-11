package lc2685

import "testing"

func Test_countCompleteComponents(t *testing.T) {
	testcases := []struct {
		n        int
		edges    [][]int
		expected int
	}{
		{6, [][]int{{0, 1}, {0, 2}, {1, 2}, {3, 4}}, 3},
		{6, [][]int{{0, 1}, {0, 2}, {1, 2}, {3, 4}, {3, 5}}, 1},
	}

	for i, tc := range testcases {
		actual := countCompleteComponents(tc.n, tc.edges)
		if actual != tc.expected {
			t.Errorf("Testcase countCompleteComponents#%02d (n=%d, edges=%v) failed: want %d, got %d",
				i+1, tc.n, tc.edges, tc.expected, actual)
		}
	}
}
