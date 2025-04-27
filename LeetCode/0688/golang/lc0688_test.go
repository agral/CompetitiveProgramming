package lc0688

import "testing"

func Test_knightProbability(t *testing.T) {
	testcases := []struct {
		n        int
		k        int
		row      int
		column   int
		expected float64
	}{
		{3, 2, 0, 0, 0.0625},
		{1, 0, 0, 0, 1.0},
	}

	for i, tc := range testcases {
		actual := knightProbability(tc.n, tc.k, tc.row, tc.column)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (n=%d, k=%d, row=%d, column=%d) failed: want %.5f, got %.5f",
				i+1, tc.n, tc.k, tc.row, tc.column, tc.expected, actual)
		}
	}
}
