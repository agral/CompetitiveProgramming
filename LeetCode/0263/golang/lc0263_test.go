package lc0263

import "testing"

func Test_isUgly(t *testing.T) {
	testcases := []struct {
		n        int
		expected bool
	}{
		{6, true},
		{1, true},
		{14, false},
	}

	for i, tc := range testcases {
		actual := isUgly(tc.n)
		if actual != tc.expected {
			t.Errorf("Testcase isUgly#%02d (%d) failed: want %t, got %t",
				i+1, tc.n, tc.expected, actual)
		}
	}
}
