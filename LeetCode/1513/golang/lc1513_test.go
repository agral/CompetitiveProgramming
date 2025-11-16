package lc1513

import "testing"

func Test_numSub(t *testing.T) {
	testcases := []struct {
		s        string
		expected int
	}{
		{"0110111", 9},
		{"101", 2},
		{"111111", 21},
	}

	for i, tc := range testcases {
		actual := numSub(tc.s)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.s, tc.expected, actual)
		}
	}
}
