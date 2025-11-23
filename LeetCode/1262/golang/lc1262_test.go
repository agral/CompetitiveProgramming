package lc1262

import "testing"

func Test_maxSumDivThree(t *testing.T) {
	testcases := []struct {
		nums     []int
		expected int
	}{
		{[]int{3, 6, 5, 1, 8}, 18},
		{[]int{4}, 0},
		{[]int{1, 2, 3, 4, 4}, 12},
	}

	for i, tc := range testcases {
		actual := maxSumDivThree(tc.nums)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.nums, tc.expected, actual)
		}
	}
}
