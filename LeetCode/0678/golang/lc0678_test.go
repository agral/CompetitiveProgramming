package lc0678

import "testing"

func Test_checkValidString(t *testing.T) {
	testcases := []struct {
		s        string
		expected bool
	}{
		{"()", true},
		{"(*)", true},
		{"(*))", true},
		{"(", false},
		{"(*)))", false},
		{"***", true},
		{"(*()*)", true}, // this can be made either "()()()", or "((()))", both ore valid.
		{"((())())", true},
	}

	for i, tc := range testcases {
		actual := checkValidString(tc.s)
		if actual != tc.expected {
			t.Errorf("Testcase checkValidString#%02d (%s) failed: want %t, got %t",
				i+1, tc.s, tc.expected, actual)
		}
	}
}
