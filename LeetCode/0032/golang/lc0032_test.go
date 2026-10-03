package lc0032

import "testing"

func Test_longestValidParentheses(t *testing.T) {
	testcases := []struct {
		s        string
		expected int
	}{
		{"(()", 2},
		{")()())", 4},
		{"", 0},
	}

	for i, tc := range testcases {
		actual := longestValidParentheses(tc.s)
		if actual != tc.expected {
			t.Errorf("Testcase longestValidParentheses#%02d (%v) failed: want %d, got %d",
				i+1, tc.s, tc.expected, actual)
		}
	}
}
