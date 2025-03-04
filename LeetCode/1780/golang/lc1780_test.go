package lc1780

import "testing"

func Test_checkPowersOfThree(t *testing.T) {
	testcases := []struct {
		n        int
		expected bool
	}{
		{12, true},
		{91, true},
		{21, false},
	}

	for i, tc := range testcases {
		actual := checkPowersOfThree(tc.n)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%d) failed: want %t, got %t",
				i+1, tc.n, tc.expected, actual)
		}
	}
}
