package lc0233

import "testing"

func Test_countDigitOne(t *testing.T) {
	testcases := []struct {
		n        int
		expected int
	}{
		{0, 0},
		{10, 2},
		{11, 4},
		{12, 5},
		{13, 6},
	}

	for i, tc := range testcases {
		actual := countDigitOne(tc.n)
		if actual != tc.expected {
			t.Errorf("Testcase countDigitOne#%02d (%v) failed: want %d, got %d",
				i+1, tc.n, tc.expected, actual)
		}
	}
}
