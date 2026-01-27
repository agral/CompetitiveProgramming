package lc3650

import "testing"

func Test_minCost(t *testing.T) {
	testcases := []struct {
		n        int
		edges    [][]int
		expected int
	}{
		{4, [][]int{{0, 1, 3}, {3, 1, 1}, {2, 3, 4}, {0, 2, 2}}, 5},
		{4, [][]int{{0, 2, 1}, {2, 1, 1}, {1, 3, 1}, {2, 3, 3}}, 3},
	}

	for i, tc := range testcases {
		actual := minCost(tc.n, tc.edges)
		if actual != tc.expected {
			t.Errorf("Testcase minCost#%02d failed: want %d, got %d",
				i+1, tc.expected, actual)
		}
	}
}
