package lc3370

import "testing"

func Test_smallestNumber(t *testing.T) {
	testcases := []struct {
		n        int
		expected int
	}{
		{1, 1},
		{2, 3},
		{3, 3},
		{4, 7},
		{5, 7},
		{6, 7},
		{7, 7},
		{8, 15},
		{15, 15},
		{254, 255},
		{255, 255},
		{256, 511},
		{512, 1023},
		{1000, 1023},
	}

	for i, tc := range testcases {
		actual := smallestNumber(tc.n)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.n, tc.expected, actual)
		}
	}
}
