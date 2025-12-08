package lc1523

import "testing"

func Test_countOdds(t *testing.T) {
	testcases := []struct {
		low      int
		high     int
		expected int
	}{
		{3, 7, 3},
		{8, 10, 1},
		{0, 1, 1},
		{0, 2, 1},
		{0, 3, 2},
		{1, 3, 2},
		{1, 4, 2},
		{1, 5, 3},
	}

	for i, tc := range testcases {
		actual := countOdds(tc.low, tc.high)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%d-%d) failed: want %d, got %d",
				i+1, tc.low, tc.high, tc.expected, actual)
		}
	}
}
