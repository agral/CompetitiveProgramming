package lc3719

import "testing"

func Test_longestBalanced(t *testing.T) {
	testcases := []struct {
		nums     []int
		expected int
	}{
		{[]int{2, 5, 4, 3}, 4},
		{[]int{3, 2, 2, 5, 4}, 5},
		{[]int{1, 2, 3, 2}, 3},
		{[]int{4, 2, 4, 1}, 2},
		{[]int{42, 42, 44, 17, 2}, 2},
		{[]int{13, 1, 12}, 2}, // WA #01
	}

	for i, tc := range testcases {
		actual := longestBalanced(tc.nums)
		if actual != tc.expected {
			t.Errorf("Testcase longestBalanced#%02d (%v) failed: want %d, got %d",
				i+1, tc.nums, tc.expected, actual)
		}
	}
}
