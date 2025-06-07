package lc3170

import "testing"

func Test_clearStars(t *testing.T) {
	testcases := []struct {
		s        string
		expected string
	}{
		{"aaba*", "aab"},
		{"abc", "abc"},
	}

	for i, tc := range testcases {
		actual := clearStars(tc.s)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%q) failed: want %q, got %q", i+1, tc.s, tc.expected, actual)
		}
	}
}
