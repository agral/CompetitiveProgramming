package lc1317

import "testing"

func Test_hasZeroDigitInDecimalRepresentation(t *testing.T) {
	testcases := []struct {
		n        int
		expected bool
	}{
		{1, false},
		{2, false},
		{10, true},
		{11, false},
		{19, false},
		{20, true},
		{99, false},
		{100, true},
		{101, true},
		{109, true},
		{110, true},
		{111, false},
		{1101, true},
		{12345678901, true},
		{12345678912, false},
	}
	for i, tc := range testcases {
		actual := hasZeroDigitInDecimalRepresentation(tc.n)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%d) failed: want %t, got %t",
				i+1, tc.n, tc.expected, actual)
		}
	}
}

func Test_getNoZeroIntegers(t *testing.T) {
	testcases := []struct {
		n int
	}{
		{2},
		{11},
	}

	for i, tc := range testcases {
		actual := getNoZeroIntegers(tc.n)
		if actual[0]+actual[1] != tc.n {
			t.Errorf("Testcase %02d (%d) failed: got [%d, %d] with sum %d != %d",
				i+1, tc.n, actual[0], actual[1], actual[0]+actual[1], tc.n)
		}
		if hasZeroDigitInDecimalRepresentation(actual[0]) {
			t.Errorf("Testcase %02d (%d) failed: got [%d, %d],"+
				" where the first component %d has digit zero in decimal representation.",
				i+1, tc.n, actual[0], actual[1], actual[0])
		}
		if hasZeroDigitInDecimalRepresentation(actual[1]) {
			t.Errorf("Testcase %02d (%d) failed: got [%d, %d],"+
				" where the second component %d has digit zero in decimal representation.",
				i+1, tc.n, actual[0], actual[1], actual[1])
		}
	}
}
