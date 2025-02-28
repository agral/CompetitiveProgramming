package lc1092

import "testing"

func Test_shortestCommonSupersequence(t *testing.T) {
	testcases := []struct {
		str1     string
		str2     string
		expected string
	}{
		{"abac", "cab", "cabac"},
		{"aaaaaaaa", "aaaaaaaa", "aaaaaaaa"}, // that's character 'a' repeated 8 times, in each string.
		{"abc", "def", "abcdef"},
	}

	for i, tc := range testcases {
		actual := shortestCommonSupersequence(tc.str1, tc.str2)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%q|%q) failed: want %q, got %q",
				i+1, tc.str1, tc.str2, tc.expected, actual)
		}
	}
}
