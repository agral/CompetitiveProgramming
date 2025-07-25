package lc3487

import "testing"

func Test_maxSum(t *testing.T) {
	testcases := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5}, 15},
		{[]int{1, 1, 0, 1, 1}, 1},
		{[]int{1, 2, -1, -2, 1, 0, -1}, 3},

		// when all the numbers are negative, the biggest one has to be returned.
		{[]int{-1, -2, -3, -4, -5}, -1},
	}

	for i, tc := range testcases {
		actual := maxSum(tc.nums)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.nums, tc.expected, actual)
		}
	}
}
