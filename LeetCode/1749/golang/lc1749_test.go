package lc1749

import "testing"

func Test_maxAbsoluteSum(t *testing.T) {
	testcases := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, -3, 2, 3, -4}, 5},     //[2, 3]
		{[]int{2, -5, 1, -4, 3, -2}, 8}, //[-5, 1, -4]
	}

	for i, tc := range testcases {
		actual := maxAbsoluteSum(tc.nums)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.nums, tc.expected, actual)
		}
	}
}
