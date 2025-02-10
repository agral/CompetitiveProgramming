package lc3174

import "testing"

func Test_clearDigits(t *testing.T) {
	testcases := []struct {
		s        string
		expected string
	}{
		{"abc", "abc"},
		{"cb34", ""},
		{"abc12", "a"},
	}

	for i, tc := range testcases {
		actual := clearDigits(tc.s)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%s) failed: want %s, got %s", i+1, tc.s, tc.expected, actual)
		}
	}
}
