package lc3010

import "testing"

func Test_minimumCost(t *testing.T) {
	testcases := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 2, 3, 12}, 6},  // {1}, {2}, {3, 12}
		{[]int{5, 4, 3}, 12},     // {5}, {4}, {3}
		{[]int{10, 3, 1, 1}, 12}, // {10, 3}, {1}, {1}
	}

	for i, tc := range testcases {
		actual := minimumCost(tc.nums)
		if actual != tc.expected {
			t.Errorf("Testcase minimumCost#%02d (%v) failed: want %d, got %d",
				i+1, tc.nums, tc.expected, actual)
		}
	}
}
