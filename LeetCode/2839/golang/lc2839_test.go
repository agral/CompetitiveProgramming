package lc2839

import "testing"

func Test_canBeEqual(t *testing.T) {
	testcases := []struct {
		s1       string
		s2       string
		expected bool
	}{
		{"abcd", "cdab", true},  // example 1
		{"abcd", "dacb", false}, // example 2

		{"wxyz", "wxyz", true},  // strings equal, no flipping
		{"wxyz", "yxwz", true},  // 0 and 2 swapped
		{"wxyz", "yzwx", true},  // both 0,2 and 1,3 swapped
		{"wxyz", "wzyx", true},  // 1 and 3 swapped
		{"wxyz", "abcd", false}, // impossible one
		{"wxyz", "zxyw", false}, // 0 and 3 swapped; should not match.
		{"wxyz", "wyxz", false}, // 1 and 2 swapped; should not match.
		{"wxyz", "zyxw", false}, // both 0,3 and 1,2 swapped; still should not match.
	}

	for i, tc := range testcases {
		actual := canBeEqual(tc.s1, tc.s2)
		if actual != tc.expected {
			t.Errorf("Testcase canBeEqual#%02d (s1=%s, s2=%s) failed: want %t, got %t",
				i+1, tc.s1, tc.s2, tc.expected, actual)
		}
	}
}
