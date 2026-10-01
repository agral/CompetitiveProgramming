package lc0020

import "testing"

func Test_isValid(t *testing.T) {
	testcases := []struct {
		s        string
		expected bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([])", true},
		{"([)]", false},
		{"(", false},
	}

	for i, tc := range testcases {
		actual := isValid(tc.s)
		if actual != tc.expected {
			t.Errorf("Testcase isValid#%02d (%s) failed: want %t, got %t",
				i+1, tc.s, tc.expected, actual)
		}
	}
}
