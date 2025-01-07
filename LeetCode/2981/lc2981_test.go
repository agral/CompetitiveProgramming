package lc2981

import "testing"

func TestMaximumLength(t *testing.T) {
	testcases := []struct {
		s        string
		expected int
	}{
		{"aaaa", 2},    // "aa" substring occurs thrice
		{"abcdef", -1}, // no substring occurs thrice
		{"abcaba", 1},  // "a" is the only substring that occurs thrice
	}

	for i, tc := range testcases {
		actual := maximumLength(tc.s)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%s) failed: want %d, got %d", i, tc.s, tc.expected, actual)
		}
	}
}
