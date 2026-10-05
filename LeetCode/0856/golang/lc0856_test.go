package lc0856

import "testing"

func Test_scoreOfParentheses(t *testing.T) {
	testcases := []struct {
		s        string
		expected int
	}{
		{"()", 1},
		{"(())", 2},
		{"()()", 2},
		{"()()()", 3},
		{"((()))", 4},
		{"((()()))", 8},
	}

	for i, tc := range testcases {
		actual := scoreOfParentheses(tc.s)
		if actual != tc.expected {
			t.Errorf("Testcase scoreOfParentheses#%02d (%s) failed: want %d, got %d",
				i+1, tc.s, tc.expected, actual)
		}
	}
}
