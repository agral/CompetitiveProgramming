package lc2116

import "testing"

func TestCanBeValid(t *testing.T) {
	testcases := []struct {
		s        string
		locked   string
		expected bool
	}{
		{"))()))", "010100", true},
		{"()()", "0000", true},
		{")", "0", false},
	}

	for i, tc := range testcases {
		actual := canBeValid(tc.s, tc.locked)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%s/%s) failed: want %v, got %v",
				i+1, tc.s, tc.locked, tc.expected, actual)
		}
	}
}
