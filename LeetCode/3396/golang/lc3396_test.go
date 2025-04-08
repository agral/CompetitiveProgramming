package lc3396

import "testing"

func Test_minimumOperations(t *testing.T) {
	testcases := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 2, 3, 4, 2, 3, 3, 5, 7}, 2},
		{[]int{4, 5, 6, 4, 4}, 2},
		{[]int{6, 7, 8, 9}, 0},
		{[]int{1}, 0},
		{[]int{1, 2}, 0},
		{[]int{1, 1}, 1},
		{[]int{1, 2, 3}, 0},
		{[]int{1, 2, 2}, 1},
		{[]int{1, 2, 3, 4}, 0},
		{[]int{1, 1, 3, 4}, 1},
		{[]int{1, 2, 3, 4, 5}, 0},
		{[]int{1, 1, 3, 4, 5}, 1},
		{[]int{1, 1, 1, 1, 1}, 2},
		{[]int{1, 1, 1, 1, 1, 1}, 2},
		{[]int{1, 1, 1, 1, 1, 1, 1}, 2},
		{[]int{1, 1, 1, 1, 1, 1, 2}, 2},
		{[]int{1, 1, 1, 1, 1, 2, 1}, 2},
		{[]int{1, 1, 1, 1, 1, 1, 1, 1}, 3},
		{[]int{1, 1, 1, 1, 1, 1, 2, 1}, 2},
		{[]int{1, 1, 1, 1, 1, 1, 1, 2}, 2},
	}

	for i, tc := range testcases {
		actual := minimumOperations(tc.nums)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.nums, tc.expected, actual)
		}
	}
}
