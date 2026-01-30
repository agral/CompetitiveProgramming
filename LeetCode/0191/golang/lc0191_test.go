package lc0191

import "testing"

func Test_hammingWeight(t *testing.T) {
	testcases := []struct {
		n        int
		expected int
	}{
		{0b1011 /*11 in decimal*/, 3},
		{0b10000000 /*128 in decimal*/, 1},
		{0b1111111111111111111111111111101 /*12147483645 in decimal*/, 30},
	}

	for i, tc := range testcases {
		actual := hammingWeight(tc.n)
		if actual != tc.expected {
			t.Errorf("Testcase hammingWeight#%02d (%d) failed: want %d, got %d",
				i+1, tc.n, tc.expected, actual)
		}
	}
}
