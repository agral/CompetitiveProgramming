package lc3300

import "testing"

func Test_possibleStringCount(t *testing.T) {
	testcases := []struct {
		word     string
		expected int
	}{
		{"abbcccc", 5},
		{"abcd", 1},
		{"aaaa", 4},
	}

	for i, tc := range testcases {
		actual := possibleStringCount(tc.word)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.word, tc.expected, actual)
		}
	}
}
