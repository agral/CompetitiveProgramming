package lc1466

import "testing"

func Test_minReorder(t *testing.T) {
	testcases := []struct {
		n           int
		connections [][]int
		expected    int
	}{
		{6, [][]int{{0, 1}, {1, 3}, {2, 3}, {4, 0}, {4, 5}}, 3},
		{5, [][]int{{1, 0}, {1, 2}, {3, 2}, {3, 4}}, 2},
		{3, [][]int{{1, 0}, {2, 0}}, 0},
	}

	for i, tc := range testcases {
		actual := minReorder(tc.n, tc.connections)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%d | %v) failed: want %d, got %d",
				i+1, tc.n, tc.connections, tc.expected, actual)
		}
	}
}
