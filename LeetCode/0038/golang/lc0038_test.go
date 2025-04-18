package lc0038

import "testing"

func Test_countAndSay(t *testing.T) {
	testcases := []struct {
		input    int
		expected string
	}{
		{1, "1"},
		{2, "11"},
		{3, "21"},
		{4, "1211"},
		{5, "111221"},
		{6, "312211"},
	}

	for i, tc := range testcases {
		actual := countAndSay(tc.input)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%d) failed: want %s, got %s", i+1, tc.input, tc.expected, actual)
		}
	}
}
