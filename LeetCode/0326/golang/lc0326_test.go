package lc0326

import "testing"

func Test_isPowerOfThree(t *testing.T) {
	testcases := []struct {
		input    int
		expected bool
	}{
		{-27, false},
		{0, false},
		{1, true}, // 1 is 3^0
		{2, false},
		{3, true}, // 3 is 3^1
		{4, false},
		{5, false},
		{6, false},
		{7, false},
		{8, false},
		{9, true}, // 9 is 3^2
		{26, false},
		{27, true},            // 27 is 3^3
		{1_162_261_467, true}, // this is 3^19.
	}

	for i, tc := range testcases {
		actual := isPowerOfThree(tc.input)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %t, got %t", i+1, tc.input, tc.expected, actual)
		}
	}
}
