package lc3021

import "testing"

func Test_flowerGame(t *testing.T) {
	testcases := []struct {
		n        int
		m        int
		expected int64
	}{
		{3, 2, int64(3)},
		{1, 1, int64(0)},
	}

	for i, tc := range testcases {
		actual := flowerGame(tc.n, tc.m)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.n, tc.expected, actual)
		}
	}
}
