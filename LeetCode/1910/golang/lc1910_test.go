package lc1910

import "testing"

func Test_removeOccurrences(t *testing.T) {
	testcases := []struct {
		s        string
		part     string
		expected string
	}{
		{"daabcbaabcbc", "abc", "dab"},
		{"axxxxyyyyb", "xy", "ab"},
	}

	for i, tc := range testcases {
		actual := removeOccurrences(tc.s, tc.part)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%q|%q) failed: want %q, got %q", i+1, tc.s, tc.part, tc.expected, actual)
		}
	}
}
