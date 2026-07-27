package lc1464

import "testing"

func Test_maxProduct(t *testing.T) {
	testcases := []struct {
		nums     []int
		expected int
	}{
		{[]int{3, 4, 5, 2}, 12},
		{[]int{1, 5, 4, 5}, 16},
		{[]int{3, 7}, 12},
		{[]int{1, 1, 1, 1, 1, 1}, 0},
		{[]int{1, 1, 2, 1, 1, 1}, 0},
		{[]int{1, 1, 2, 1, 2, 1}, 1},
	}

	for i, tc := range testcases {
		actual := maxProduct(tc.nums)
		if actual != tc.expected {
			t.Errorf("Testcase maxProduct#%02d (%v) failed: want %d, got %d",
				i+1, tc.nums, tc.expected, actual)
		}
	}
}
