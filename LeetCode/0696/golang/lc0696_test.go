package lc0696

import "testing"

func Test_countBinarySubstrings(t *testing.T) {
	testcases := []struct {
		s        string
		expected int
	}{
		{"00110011", 6},
		{"10101", 4},
		{"111", 0},
		{"0000", 0},
		{"1000", 1},
		{"0001", 1},
		{"0011000111", 7}, //0011, 1100, 000111 -> 7 in total.
	}

	for i, tc := range testcases {
		actual := countBinarySubstrings(tc.s)
		if actual != tc.expected {
			t.Errorf("Testcase countBinarySubstrings#%02d (%v) failed: want %d, got %d",
				i+1, tc.s, tc.expected, actual)
		}
	}
}
