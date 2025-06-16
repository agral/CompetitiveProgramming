package lc2016

import "testing"

func Test_maximumDifference(t *testing.T) {
	testcases := []struct {
		nums     []int
		expected int
	}{
		{[]int{7, 1, 5, 4}, 4},
		{[]int{9, 4, 3, 2}, -1},
		{[]int{1, 5, 2, 10}, 9},
		{[]int{1, 1, 1, 1}, -1},
		{[]int{5, 1, 2, 1}, 1},
	}

	for i, tc := range testcases {
		actual := maximumDifference(tc.nums)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.nums, tc.expected, actual)
		}
	}
}
