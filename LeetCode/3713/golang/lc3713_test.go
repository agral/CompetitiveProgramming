package lc3713

import "testing"

func Test_longestBalanced(t *testing.T) {
	testcases := []struct {
		s        string
		expected int
	}{
		{"abbac", 4},   // "abba"
		{"zzabccy", 4}, // "zabc"
		{"aba", 2},     // "ab" or "ba"
		{"z", 1},
		{"aaaaaa", 6},
		{"aabaaa", 3}, // note: only the last "aaa" matches.
	}

	for i, tc := range testcases {
		actual := longestBalanced(tc.s)
		if actual != tc.expected {
			t.Errorf("Testcase longestBalanced#%02d (%s) failed: want %d, got %d",
				i+1, tc.s, tc.expected, actual)
		}
	}
}
