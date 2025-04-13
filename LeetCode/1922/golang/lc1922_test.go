package lc1922

import "testing"

func Test_countGoodNumbers(t *testing.T) {
	testcases := []struct {
		n        int64
		expected int
	}{
		{1, 5},
		{4, 400},
		{50, 564908303},
	}

	for i, tc := range testcases {
		actual := countGoodNumbers(tc.n)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%d) failed: want %d, got %d", i+1, tc.n, tc.expected, actual)
		}
	}
}
