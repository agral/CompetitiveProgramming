package lc1800

import "testing"

func Test_maxAscendingSum(t *testing.T) {
	testcases := []struct {
		nums     []int
		expected int
	}{
		{[]int{10, 20, 30, 5, 10, 50}, 65},
		{[]int{10, 20, 30, 40, 50}, 150},
		{[]int{12, 17, 15, 13, 10, 11, 12}, 33},
		{[]int{13, 13, 13}, 13},
		{[]int{13, 12, 11}, 13},
	}

	for i, tc := range testcases {
		actual := maxAscendingSum(tc.nums)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.nums, tc.expected, actual)
		}
	}
}
