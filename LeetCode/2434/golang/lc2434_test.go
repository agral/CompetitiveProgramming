package lc2434

import "testing"

func Test_robotWithString(t *testing.T) {
	testcases := []struct {
		input    string
		expected string
	}{
		{"zza", "azz"},
		{"bac", "abc"},
		{"bdda", "addb"},
	}

	for i, tc := range testcases {
		actual := robotWithString(tc.input)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%q) failed: want %q, got %q", i+1, tc.input, tc.expected, actual)
		}
	}
}
