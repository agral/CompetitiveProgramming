package lc1009

import "testing"

func Test_bitwiseComplement(t *testing.T) {
	testcases := []struct {
		n        int
		expected int
	}{
		{5, 2},  // 101 -> 010
		{7, 0},  // 111 -> 000
		{10, 5}, // 1010 -> 0101
	}

	for i, tc := range testcases {
		actual := bitwiseComplement(tc.n)
		if actual != tc.expected {
			t.Errorf("Testcase bitwiseComplement#%02d (%v) failed: want %d, got %d",
				i+1, tc.n, tc.expected, actual)
		}
	}
}
