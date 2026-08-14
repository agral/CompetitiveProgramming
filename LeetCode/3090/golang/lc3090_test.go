package lc3090

import "testing"

func Test_maximumLengthSubstring(t *testing.T) {
	testcases := []struct {
		s        string
		expected int
	}{
		{"bcbbbcba", 4}, // the last "bcba" fits.
		{"aaaa", 2},
		{"abba", 4},     // yay ABBA
		{"abbba", 3},    // both "abb" and "bba" do fit, both are of length 3.
		{"abababab", 4}, // any "abab" / "baba" substrings fit.
		{"abcdefghijklmnopqrstuvwxyz", 'z' - 'a' + 1},
	}

	for i, tc := range testcases {
		actual := maximumLengthSubstring(tc.s)
		if actual != tc.expected {
			t.Errorf("Testcase maximumLengthSubstring#%02d (%v) failed: want %d, got %d",
				i+1, tc.s, tc.expected, actual)
		}
	}
}
