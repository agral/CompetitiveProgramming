package lc0628

import "testing"

func Test_maximumProduct(t *testing.T) {
	testcases := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 2, 3}, 6},
		{[]int{1, 2, 3, 4}, 24},
		{[]int{-1, -2, -3}, -6},
	}

	for i, tc := range testcases {
		actual := maximumProduct(tc.nums)
		if actual != tc.expected {
			t.Errorf("Testcase maximumProduct#%02d (%v) failed: want %d, got %d",
				i+1, tc.nums, tc.expected, actual)
		}
	}
}
