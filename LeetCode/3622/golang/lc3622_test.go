package lc3622

import "testing"

func Test_checkDivisibility(t *testing.T) {
	testcases := []struct {
		n        int
		expected bool
	}{
		{99, true},
		{23, false},
		{10, true},
		{11, false},
	}

	for i, tc := range testcases {
		actual := checkDivisibility(tc.n)
		if actual != tc.expected {
			t.Errorf("Testcase checkDivisibility#%02d (%v) failed: want %t, got %t",
				i+1, tc.n, tc.expected, actual)
		}
	}
}
