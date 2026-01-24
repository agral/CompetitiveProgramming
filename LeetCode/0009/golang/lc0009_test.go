package lc0009

import "testing"

func Test_isPalindrome(t *testing.T) {
	testcases := []struct {
		x        int
		expected bool
	}{
		{121, true},
		{-121, false},
		{10, false},
	}

	for i, tc := range testcases {
		actual := isPalindrome(tc.x)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%d) failed: want %t, got %t", i+1, tc.x, tc.expected, actual)
		}
	}
}
