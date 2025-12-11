package lc3531

import "testing"

func Test_countCoveredBuildings(t *testing.T) {
	testcases := []struct {
		n         int
		buildings [][]int
		expected  int
	}{
		{3, [][]int{{1, 2}, {2, 2}, {3, 2}, {2, 1}, {2, 3}}, 1},
		{3, [][]int{{1, 1}, {1, 2}, {2, 1}, {2, 2}}, 0},
		{5, [][]int{{1, 3}, {3, 2}, {3, 3}, {3, 5}, {5, 3}}, 1},
	}

	for i, tc := range testcases {
		actual := countCoveredBuildings(tc.n, tc.buildings)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v | %v) failed: want %d, got %d",
				i+1, tc.n, tc.buildings, tc.expected, actual)
		}
	}
}
