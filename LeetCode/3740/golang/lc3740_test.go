package lc3740

import "testing"

func Test_minimumDistance(t *testing.T) {
	testcases := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 2, 1, 1, 3}, 6},
		{[]int{1, 1, 2, 3, 2, 1, 2}, 8},
		{[]int{1}, -1},
	}

	for i, tc := range testcases {
		actual := minimumDistance(tc.nums)
		if actual != tc.expected {
			t.Errorf("Testcase minimumDistance#%02d (%v) failed: want %d, got %d",
				i+1, tc.nums, tc.expected, actual)
		}
	}
}
