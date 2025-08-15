package lc0342

import "testing"

func Test_isPowerOfFour(t *testing.T) {
	testcases := []struct {
		n        int
		expected bool
	}{
		{1, true},
		{2, false},
		{3, false},
		{4, true},
		{8, false},
		{9, false},
		{16, true},
		{64, true},
		{256, true},
		{1024, true},
		{8192, false},
	}

	for i, tc := range testcases {
		actual := isPowerOfFour(tc.n)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%d) failed: want %t, got %t", i+1, tc.n, tc.expected, actual)
		}
	}
}
