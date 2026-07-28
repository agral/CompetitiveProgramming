package lc3517

import "testing"

func Test_smallestPalindrome(t *testing.T) {
	testcases := []struct {
		s        string
		expected string
	}{
		{"z", "z"},
		{"babab", "abbba"},
		{"daccad", "acddca"},
	}

	for i, tc := range testcases {
		actual := smallestPalindrome(tc.s)
		if actual != tc.expected {
			t.Errorf("Testcase smallestPalindrome#%02d (%s) failed: want %s, got %s",
				i+1, tc.s, tc.expected, actual)
		}
	}
}
