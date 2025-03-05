package lc2579

import "testing"

func Test_coloredCells(t *testing.T) {
	testcases := []struct {
		n        int
		expected int64
	}{
		{1, 1},
		{2, 5},
		{3, 13},
		{4, 25},
	}

	for i, tc := range testcases {
		actual := coloredCells(tc.n)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.n, tc.expected, actual)
		}
	}
}
