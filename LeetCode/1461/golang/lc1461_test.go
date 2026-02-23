package lc1461

import "testing"

func Test_hasAllCodes(t *testing.T) {
	testcases := []struct {
		s        string
		k        int
		expected bool
	}{
		{"00110110", 2, true},
		{"0110", 1, true},
		{"0110", 2, false},
	}

	for i, tc := range testcases {
		actual := hasAllCodes(tc.s, tc.k)
		if actual != tc.expected {
			t.Errorf("Testcase hasAllCodes#%02d (%s | %d) failed: want %t, got %t",
				i+1, tc.s, tc.k, tc.expected, actual)
		}
	}
}
