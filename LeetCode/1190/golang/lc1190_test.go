package lc1190

import "testing"

func Test_reverseParentheses(t *testing.T) {
	testcases := []struct {
		s        string
		expected string
	}{
		{"(abcd)", "dcba"},
		{"(u(love)i)", "iloveu"},
		{"(ed(et(oc))el)", "leetcode"},
		{"(rg)(la)", "gral"},
	}

	for i, tc := range testcases {
		actual := reverseParentheses(tc.s)
		if actual != tc.expected {
			t.Errorf("Testcase reverseParentheses#%02d (%s) failed: want %s, got %s",
				i+1, tc.s, tc.expected, actual)
		}
	}
}
