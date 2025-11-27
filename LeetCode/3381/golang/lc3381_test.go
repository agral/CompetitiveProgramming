package lc3381

import "testing"

func Test_maxSubarraySum(t *testing.T) {
	testcases := []struct {
		nums     []int
		k        int
		expected int64
	}{
		{[]int{1, 2}, 1, int64(3)},
		{[]int{-1, -2, -3, -4, -5}, 4, int64(-10)},
		{[]int{-5, 1, 2, -3, 4}, 2, int64(4)},
	}

	for i, tc := range testcases {
		actual := maxSubarraySum(tc.nums, tc.k)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v | %d) failed: want %d, got %d",
				i+1, tc.nums, tc.k, tc.expected, actual)
		}
	}
}
