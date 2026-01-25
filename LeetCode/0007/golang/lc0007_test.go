package lc0007

import "testing"

func Test_reverse(t *testing.T) {
	testcases := []struct {
		input    int
		expected int
	}{
		{123, 321},
		{-123, -321},
		{120, 21},
	}

	for i, tc := range testcases {
		actual := reverse(tc.input)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.input, tc.expected, actual)
		}
	}
}
