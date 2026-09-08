package lc3870

import "testing"

func Test_countCommas(t *testing.T) {
	testcases := []struct {
		n        int
		expected int
	}{
		{1_002, 3}, // 1_000, 1_001, 1_002
		{998, 0},
		{100_000, 99_001}, // ==100_000 - 999
	}

	for i, tc := range testcases {
		actual := countCommas(tc.n)
		if actual != tc.expected {
			t.Errorf("Testcase countCommas#%02d (%v) failed: want %d, got %d",
				i+1, tc.n, tc.expected, actual)
		}
	}
}
