package lc1957

import "testing"

func Test_makeFancyString(t *testing.T) {
	testcases := []struct {
		s        string
		expected string
	}{
		{"leeetcode", "leetcode"},
		{"aaabaaaa", "aabaa"},
		{"aab", "aab"},
		{"aabbcc", "aabbcc"},
		{"aaaaaaaaaaa", "aa"},
	}

	for i, tc := range testcases {
		actual := makeFancyString(tc.s)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%s) failed: want %s, got %s", i+1, tc.s, tc.expected, actual)
		}
	}
}
