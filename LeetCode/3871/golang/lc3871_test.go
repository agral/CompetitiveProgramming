package lc3871

import "testing"

func Test_countCommas(t *testing.T) {
	testcases := []struct {
		n        int64
		expected int64
	}{
		{1_002, 3}, // 1_000, 1_001, 1_002
		{998, 0},
		{100_000, 99_001}, // ==100_000 - 999
		{1_000_000_000_000_000, 3_998_998_998_999_005}, // the maximum allowed number in the constraints
	}

	for i, tc := range testcases {
		actual := countCommas(tc.n)
		if actual != tc.expected {
			t.Errorf("Testcase countCommas#%02d (%v) failed: want %d, got %d",
				i+1, tc.n, tc.expected, actual)
		}
	}
}
