package lc1015

import "testing"

func Test_smallestRepunitDivByK(t *testing.T) {
	testcases := []struct {
		k        int
		expected int
	}{
		{1, 1},
		{2, -1},
		{3, 3},
		{4, -1},
		{5, -1},
		{9, 9},
	}

	for i, tc := range testcases {
		actual := smallestRepunitDivByK(tc.k)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.k, tc.expected, actual)
		}
	}
}
