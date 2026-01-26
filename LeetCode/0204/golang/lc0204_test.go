package lc0204

import "testing"

func Test_countPrimes(t *testing.T) {
	testcases := []struct {
		n        int
		expected int
	}{
		{10, 4}, // 2, 3, 5, 7
		{0, 0},  // no primes there
		{1, 0},  // nor here
	}

	for i, tc := range testcases {
		actual := countPrimes(tc.n)
		if actual != tc.expected {
			t.Errorf("Testcase countPrimes#%02d (%v) failed: want %d, got %d",
				i+1, tc.n, tc.expected, actual)
		}
	}
}
