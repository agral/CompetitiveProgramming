package lc2131

import "testing"

func Test_longestPalindrome(t *testing.T) {
	testcases := []struct {
		words    []string
		expected int
	}{
		{[]string{"lc", "cl", "gg"}, 6},                   // e.g. "clgglc"
		{[]string{"ab", "ty", "yt", "lc", "cl", "ab"}, 8}, // e.g. "cltyytlc"
		{[]string{"cc", "ll", "xx"}, 2},                   // e.g. "cc"
	}

	for i, tc := range testcases {
		actual := longestPalindrome(tc.words)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.words, tc.expected, actual)
		}
	}
}
