package lc3234

import "testing"

func Test_numberOfSubstrings(t *testing.T) {
	testcases := []struct {
		s        string
		expected int
	}{
		{"00011", 5},
		{"101101", 16},
	}

	for i, tc := range testcases {
		actual := numberOfSubstrings(tc.s)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.s, tc.expected, actual)
		}
	}
}
