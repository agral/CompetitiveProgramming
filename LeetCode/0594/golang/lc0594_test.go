package lc0594

import "testing"

func Test_findLHS(t *testing.T) {
	testcases := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 3, 2, 2, 5, 2, 3, 7}, 5}, // {3, 2, 2, 2, 3}
		{[]int{1, 2, 3, 4}, 2},             // either of: {1, 2}, or {2, 3}, or {3, 4}
		{[]int{1, 1, 1, 1}, 0},             // no harmonious subsequences at all.
	}

	for i, tc := range testcases {
		actual := findLHS(tc.nums)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.nums, tc.expected, actual)
		}
	}
}
