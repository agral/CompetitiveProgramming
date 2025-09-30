package lc2221

import "testing"

func Test_triangularSum(t *testing.T) {
	testcases := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5}, 8},
		{[]int{5}, 5},
		{[]int{3, 7}, 0},
		{[]int{9, 9, 9, 9, 9}, 4},
		{[]int{1, 1, 1, 1, 1}, 6},
	}

	for i, tc := range testcases {
		actual := triangularSum(tc.nums)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.nums, tc.expected, actual)
		}
	}
}
