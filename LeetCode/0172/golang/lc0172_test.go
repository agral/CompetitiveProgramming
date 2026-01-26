package lc0172

import "testing"

func Test_trailingZeroes(t *testing.T) {
	testcases := []struct {
		n        int
		expected int
	}{
		{0, 0},
		{3, 0},
		{5, 1},
		{10, 2}, // 10! == 3_628_800, two trailing zeroes.
		{25, 6}, // six trailing zeroes.
	}

	for i, tc := range testcases {
		actual := trailingZeroes(tc.n)
		if actual != tc.expected {
			t.Errorf("Testcase trailingZeroes#%02d (%v) failed: want %d, got %d",
				i+1, tc.n, tc.expected, actual)
		}
	}
}
