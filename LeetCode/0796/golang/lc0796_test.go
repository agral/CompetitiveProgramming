package lc0796

import "testing"

func Test_rotateString(t *testing.T) {
	testcases := []struct {
		s        string
		goal     string
		expected bool
	}{
		{"abcde", "cdeab", true},
		{"abcde", "abced", false},
		{"aa", "a", false},
	}

	for i, tc := range testcases {
		actual := rotateString(tc.s, tc.goal)
		if actual != tc.expected {
			t.Errorf("Testcase rotateString#%02d (%v) failed: want %t, got %t",
				i+1, tc.s, tc.expected, actual)
		}
	}
}
