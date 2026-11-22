package lc2124

import "testing"

func Test_checkString(t *testing.T) {
	testcases := []struct {
		s        string
		expected bool
	}{
		{"aaabbb", true},
		{"abab", false},
		{"bbb", true},
	}

	for i, tc := range testcases {
		actual := checkString(tc.s)
		if actual != tc.expected {
			t.Errorf("Testcase checkString#%02d (%s) failed: want %t, got %t",
				i+1, tc.s, tc.expected, actual)
		}
	}
}
