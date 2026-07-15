package lc3658

import "testing"

func Test_gcdOfOddEvenSums(t *testing.T) {
	testcases := []struct {
		n        int
		expected int
	}{
		{4, 4},
		{5, 5},
		{1000, 1000},
	}

	for i, tc := range testcases {
		actual := gcdOfOddEvenSums(tc.n)
		if actual != tc.expected {
			t.Errorf("Testcase gcdOfOddEvenSums#%02d (%v) failed: want %d, got %d",
				i+1, tc.n, tc.expected, actual)
		}
	}
}
