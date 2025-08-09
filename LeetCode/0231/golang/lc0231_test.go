package lc0231

import "testing"

func Test_isPowerOfTwo(t *testing.T) {
	testcases := []struct {
		n        int
		expected bool
	}{
		{1, true},
		{16, true},
		{3, false},
	}

	for i, tc := range testcases {
		actual := isPowerOfTwo(tc.n)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %t, got %t", i+1, tc.n, tc.expected, actual)
		}
	}
}
