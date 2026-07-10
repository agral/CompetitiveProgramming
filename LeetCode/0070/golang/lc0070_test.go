package lc0070

import "testing"

func Test_climbStairs(t *testing.T) {
	testcases := []struct {
		n        int
		expected int
	}{
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 5},
		{5, 8},
		{6, 13},
		{7, 21},
		{8, 34},
		{9, 55},
		{10, 89},
		{45, 1836311903},
	}

	for i, tc := range testcases {
		actual := climbStairs(tc.n)
		if actual != tc.expected {
			t.Errorf("Testcase climbStairs#%02d (%v) failed: want %d, got %d",
				i+1, tc.n, tc.expected, actual)
		}
	}
}
