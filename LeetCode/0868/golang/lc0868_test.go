package lc0868

import "testing"

func Test_binaryGap(t *testing.T) {
	testcases := []struct {
		n        int
		expected int
	}{
		{22, 2},
		{8, 0},
		{5, 2},
		{9, 3},  // 1001 -> 3
		{37, 3}, // 100101 -> 3
		{41, 3}, // 101001 -> 3
		{42, 2}, // 101010 -> 2
		{63, 1}, // 111111 -> 1
		{64, 0}, // 1000000 -> 0
	}

	for i, tc := range testcases {
		actual := binaryGap(tc.n)
		if actual != tc.expected {
			t.Errorf("Testcase binaryGap#%02d (%v) failed: want %d, got %d",
				i+1, tc.n, tc.expected, actual)
		}
	}
}
