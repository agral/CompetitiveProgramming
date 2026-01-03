package lc1411

import "testing"

func Test_numOfWays(t *testing.T) {
	testcases := []struct {
		n        int
		expected int
	}{
		{1, 12},
		{5000, 30228214},
	}

	for i, tc := range testcases {
		actual := numOfWays(tc.n)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.n, tc.expected, actual)
		}
	}
}
