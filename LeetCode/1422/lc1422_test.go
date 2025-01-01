package lc1422

import "testing"

func TestMaxScore(t *testing.T) {
	type testcase struct {
		input    string
		expected int
	}
	testcases := []testcase{
		{"011101", 5}, // left: 0 (1), right: 11101 (4)
		{"00111", 5},  // left: 00 (2), right: 111 (3)
		{"1111", 3},   // left: 1 (0), right: 111 (3)
		{"0000", 3},   // left: 000 (3), right: 0 (0)
		{"10", 0},     // left: 1 (0), right: 0 (0); this is the only possible division.
		{"00", 1},     // left: 0(1), right: 0 (0); this is the only possible division.
	}
	for i, tc := range testcases {
		actual := maxScore(tc.input)
		if actual != tc.expected {
			t.Errorf("Testcase #%d: input %s, got %d, want %d", i, tc.input, actual, tc.expected)
		}
	}
}
